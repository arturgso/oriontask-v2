<script setup lang="ts">
import { ref } from 'vue';
import { store } from '../../store';

const emit = defineEmits(['close']);

const name = ref('');
const isLoading = ref(false);
const error = ref('');

const submit = async () => {
  if (!name.value.trim()) return;

  isLoading.value = true;
  error.value = '';
  
  try {
    await store.addDharma(name.value.trim());
    emit('close');
  } catch (e: any) {
    error.value = e.message || 'Failed to create Dharma';
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal-content">
      <h2 class="modal-header">New Dharma</h2>
      
      <div v-if="error" class="bg-red-900/20 text-red-400 p-3 rounded mb-4 text-sm border border-red-900/50">
        {{ error }}
      </div>

      <div class="form-group">
        <label>Dharma Name *</label>
        <input 
          v-model="name" 
          type="text" 
          class="form-control" 
          placeholder="e.g. Work, Health, Personal Development" 
          autofocus
          @keyup.enter="submit"
          maxlength="60"
        >
      </div>

      <div class="modal-footer">
        <button class="btn" @click="emit('close')" :disabled="isLoading">Cancel</button>
        <button 
          class="btn btn-primary" 
          @click="submit" 
          :disabled="!name.trim() || isLoading"
        >
          {{ isLoading ? 'Creating...' : 'Create Dharma' }}
        </button>
      </div>
    </div>
  </div>
</template>
