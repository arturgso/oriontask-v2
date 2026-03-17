<script setup lang="ts">
import { ref, computed } from 'vue';
import { store } from '../../store';

const emit = defineEmits(['close']);

const title = ref('');
const dharmaId = ref('');
const projectId = ref('');
const milestoneId = ref('');

const dharmas = computed(() => store.dharmas);
const projects = computed(() => {
  if (!dharmaId.value) return [];
  return store.projects.filter(p => p.dharma_id === dharmaId.value);
});
const milestones = computed(() => {
  if (!projectId.value) return [];
  return store.milestones.filter(m => m.project_id === projectId.value);
});

const submit = () => {
  if (!title.value.trim() || !dharmaId.value) return;

  store.addTask({
    title: title.value.trim(),
    description: '',
    dharma_id: dharmaId.value,
    project_id: projectId.value || undefined,
    milestone_id: milestoneId.value || undefined
  });
  
  emit('close');
};
</script>

<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal-content">
      <h2 class="modal-header">New Task</h2>
      
      <div class="form-group">
        <label>Task Title *</label>
        <input v-model="title" type="text" class="form-control" placeholder="What needs to be done?" autofocus>
      </div>

      <div class="form-group">
        <label>Dharma *</label>
        <select v-model="dharmaId" class="form-control" @change="projectId = ''; milestoneId = ''">
          <option disabled value="">Select a Dharma</option>
          <option v-for="d in dharmas" :key="d.id" :value="d.id">{{ d.name }}</option>
        </select>
      </div>

      <div class="form-group" v-if="dharmaId && projects.length > 0">
        <label>Project (Optional)</label>
        <select v-model="projectId" class="form-control" @change="milestoneId = ''">
          <option value="">None</option>
          <option v-for="p in projects" :key="p.id" :value="p.id">{{ p.title }}</option>
        </select>
      </div>

      <div class="form-group" v-if="projectId && milestones.length > 0">
        <label>Milestone (Optional)</label>
        <select v-model="milestoneId" class="form-control">
          <option value="">None</option>
          <option v-for="m in milestones" :key="m.id" :value="m.id">{{ m.title }}</option>
        </select>
      </div>

      <div class="modal-footer">
        <button class="btn" @click="emit('close')">Cancel</button>
        <button class="btn btn-primary" @click="submit" :disabled="!title.trim() || !dharmaId">Create Task</button>
      </div>
    </div>
  </div>
</template>
