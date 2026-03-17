<script setup lang="ts">
import { ref, computed } from 'vue';
import { store } from '../../store';
import CustomSelect from '../ui/CustomSelect.vue';

const emit = defineEmits(['close']);

const title = ref('');
const dharmaId = ref('');
const projectId = ref('');
const milestoneId = ref('');

const dharmas = computed(() => store.dharmas.map(d => ({ id: d.id, name: d.name })));
const projects = computed(() => {
  if (!dharmaId.value) return [];
  return store.projects
    .filter(p => p.dharma_id === dharmaId.value)
    .map(p => ({ id: p.id, name: p.title }));
});
const milestones = computed(() => {
  if (!projectId.value) return [];
  return store.milestones
    .filter(m => m.project_id === projectId.value)
    .map(m => ({ id: m.id, name: m.title }));
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
        <CustomSelect 
          v-model="dharmaId"
          :options="dharmas"
          label="Dharma *"
          placeholder="Select a Dharma"
          @update:modelValue="projectId = ''; milestoneId = ''"
        />
      </div>

      <div class="form-group" v-if="dharmaId && projects.length > 0">
        <CustomSelect 
          v-model="projectId"
          :options="[{ id: '', name: 'None' }, ...projects]"
          label="Project (Optional)"
          placeholder="Select a Project"
          @update:modelValue="milestoneId = ''"
        />
      </div>

      <div class="form-group" v-if="projectId && milestones.length > 0">
        <CustomSelect 
          v-model="milestoneId"
          :options="[{ id: '', name: 'None' }, ...milestones]"
          label="Milestone (Optional)"
          placeholder="Select a Milestone"
        />
      </div>

      <div class="modal-footer">
        <button class="btn" @click="emit('close')">Cancel</button>
        <button class="btn btn-primary" @click="submit" :disabled="!title.trim() || !dharmaId">Create Task</button>
      </div>
    </div>
  </div>
</template>
