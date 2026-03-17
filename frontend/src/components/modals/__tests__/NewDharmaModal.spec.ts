import { describe, it, expect, vi, beforeEach } from 'vitest';
import { mount } from '@vue/test-utils';
import NewDharmaModal from '../NewDharmaModal.vue';
import { store } from '../../../store';

// Mock the store to avoid actual Wails bridge calls during tests
vi.mock('../../../store', () => ({
  store: {
    addDharma: vi.fn(),
  },
}));

describe('NewDharmaModal.vue', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders correctly', () => {
    const wrapper = mount(NewDharmaModal);
    expect(wrapper.find('.modal-header').text()).toBe('New Dharma');
    expect(wrapper.find('input').exists()).toBe(true);
  });

  it('disables the create button when input is empty', async () => {
    const wrapper = mount(NewDharmaModal);
    const button = wrapper.find('.btn-primary');
    
    // Check initial state
    expect((button.element as HTMLButtonElement).disabled).toBe(true);
    
    // Set value and check again
    await wrapper.find('input').setValue('Work');
    expect((button.element as HTMLButtonElement).disabled).toBe(false);
  });

  it('calls store.addDharma and emits close on success', async () => {
    const wrapper = mount(NewDharmaModal);
    const input = wrapper.find('input');
    
    await input.setValue('Health');
    await wrapper.find('.btn-primary').trigger('click');

    expect(store.addDharma).toHaveBeenCalledWith('Health');
    
    // Since submit is async, we wait for the next tick or use vi.waitFor
    await vi.waitFor(() => {
        expect(wrapper.emitted()).toHaveProperty('close');
    });
  });

  it('displays error message when store.addDharma fails', async () => {
    const errorMessage = 'Validation failed';
    (store.addDharma as any).mockRejectedValueOnce(new Error(errorMessage));
    
    const wrapper = mount(NewDharmaModal);
    await wrapper.find('input').setValue('Invalid');
    await wrapper.find('.btn-primary').trigger('click');

    await vi.waitFor(() => {
        expect(wrapper.find('.text-red-400').text()).toContain(errorMessage);
        expect(wrapper.emitted('close')).toBeUndefined();
    });
  });
});
