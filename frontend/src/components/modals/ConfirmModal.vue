<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  title: string;
  message: string;
  confirmLabel?: string;
  cancelLabel?: string;
  isDestructive?: boolean;
}>();

const emit = defineEmits(['confirm', 'close']);

const isLoading = ref(false);

const handleConfirm = async () => {
  isLoading.value = true;
  try {
    emit('confirm');
  } catch (e) {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal-content max-w-[400px]">
      <h2 class="modal-header">{{ title }}</h2>
      
      <p class="mb-6 text-text-secondary leading-relaxed">
        {{ message }}
      </p>

      <div class="modal-footer">
        <button class="btn" @click="emit('close')" :disabled="isLoading">
          {{ cancelLabel || 'Cancel' }}
        </button>
        <button 
          class="btn" 
          :class="isDestructive ? 'bg-red-500 hover:bg-red-600' : 'btn-primary'"
          @click="handleConfirm" 
          :disabled="isLoading"
        >
          {{ isLoading ? 'Processing...' : (confirmLabel || 'Confirm') }}
        </button>
      </div>
    </div>
  </div>
</template>
