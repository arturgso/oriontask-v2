<script setup lang="ts">
import { ref } from 'vue';
import { store } from '../../store';

const emit = defineEmits(['close']);

const type = ref<'positive' | 'negative'>('positive');
const description = ref('');
const date = ref(new Date().toISOString().split('T')[0]);

const submit = () => {
  if (!description.value.trim()) return;

  // Convert local date back to ISO string for storage
  const isoDate = new Date(date.value).toISOString();

  store.addKarma({
    type: type.value,
    description: description.value.trim()
  }, isoDate);
  
  emit('close');
};
</script>

<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal-content">
      <h2 class="modal-header">Log Karma</h2>
      
      <div class="form-group">
        <label>Type</label>
        <div class="flex gap-4">
          <label class="flex items-center gap-2 cursor-pointer">
            <input type="radio" v-model="type" value="positive">
            Positive
          </label>
          <label class="flex items-center gap-2 cursor-pointer">
            <input type="radio" v-model="type" value="negative">
            Negative
          </label>
        </div>
      </div>

      <div class="form-group">
        <label>Description *</label>
        <input v-model="description" type="text" class="form-control" placeholder="What happened?" autofocus>
      </div>

      <div class="form-group">
        <label>Date</label>
        <input v-model="date" type="date" class="form-control">
      </div>

      <div class="modal-footer">
        <button class="btn" @click="emit('close')">Cancel</button>
        <button class="btn btn-primary" @click="submit" :disabled="!description.trim()">Save Entry</button>
      </div>
    </div>
  </div>
</template>
