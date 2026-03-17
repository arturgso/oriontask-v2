<script setup lang="ts">
import { computed } from 'vue';
import { store } from '../../store';
import { Circle, ArrowRight, Plus } from 'lucide-vue-next';

// Computed properties to derive active tasks. Limit to 5.
const activeTasks = computed(() => {
  return store.tasks.filter(t => t.status === 'active').slice(0, 5);
});

const completeTask = (id: string) => {
  store.completeTask(id);
};

const getDharmaName = (dharmaId: string) => {
  return store.dharmas.find(d => d.id === dharmaId)?.name || 'General';
};

const getProjectName = (projectId?: string) => {
  if (!projectId) return null;
  return store.projects.find(p => p.id === projectId)?.title || null;
};

const getMetaText = (task: any) => {
  const dharma = getDharmaName(task.dharma_id);
  const project = getProjectName(task.project_id);
  return project ? `${dharma} · ${project}` : dharma;
};
</script>

<template>
  <div class="pt-[100px] flex flex-col w-full">
    <h1 class="font-serif text-[32px] mb-12 font-normal text-text-primary opacity-90 tracking-[-0.01em]">Now</h1>
    
    <div v-if="activeTasks.length === 0" class="py-10 text-text-secondary text-left">
      <p>No active tasks. You're all caught up!</p>
    </div>
    
    <div class="flex flex-col">
      <div v-for="(task, index) in activeTasks" :key="task.id" class="flex items-start gap-4 py-5 border-b border-border transition-all duration-200 first:border-t">
        <button class="bg-transparent border-none text-text-secondary cursor-pointer p-0 mt-1 flex items-center justify-center transition-all duration-200 opacity-40 hover:text-success hover:opacity-100 hover:scale-110" @click="completeTask(task.id)">
          <Circle :size="24" class="stroke-[1px]" />
        </button>
        
        <div class="flex-1 flex justify-between items-center">
          <div class="flex flex-col gap-1">
            <h3 class="m-0 text-base font-normal text-text-primary tracking-[-0.01em]">{{ task.title }}</h3>
            <span class="text-[13px] text-text-secondary opacity-70">{{ getMetaText(task) }}</span>
          </div>
          <ArrowRight v-if="index === 0" :size="18" class="text-text-secondary opacity-20" />
        </div>
      </div>
      
      <button class="flex items-center gap-2 bg-transparent border-none text-text-secondary text-sm py-8 cursor-pointer transition-all duration-200 opacity-80 hover:text-text-primary hover:opacity-100" @click="$emit('openNewTaskModal')">
        <Plus :size="18" />
        <span>Add task</span>
      </button>
    </div>
  </div>
</template>
