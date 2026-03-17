<script setup lang="ts">
import { computed } from 'vue';
import { store } from '../../store';
import { Check } from 'lucide-vue-next';

const groupedTasks = computed(() => {
  const groups: Record<string, typeof store.tasks> = {};
  
  const completed = store.tasks.filter(t => t.status === 'completed');
  const sorted = [...completed].sort((a, b) => {
    const timeA = a.completed_at ? new Date(a.completed_at).getTime() : 0;
    const timeB = b.completed_at ? new Date(b.completed_at).getTime() : 0;
    return timeB - timeA;
  });

  sorted.forEach(task => {
    const dateStr = task.completed_at 
      ? new Date(task.completed_at).toLocaleDateString(undefined, { month: 'long', day: 'numeric', year: 'numeric' })
      : 'Unknown Date';
    if (!groups[dateStr]) groups[dateStr] = [];
    groups[dateStr].push(task);
  });

  return groups;
});

const getDharmaName = (dharmaId: string) => {
  return store.dharmas.find(d => d.id === dharmaId)?.name || 'Unknown';
};
const getProjectName = (projectId?: string) => {
  if (!projectId) return null;
  return store.projects.find(p => p.id === projectId)?.title || null;
};
</script>

<template>
  <div class="pt-[80px] flex flex-col mb-20 w-full">
    <!-- Header -->
    <header class="flex justify-between items-start mb-12">
      <div>
        <h1 class="font-serif text-[32px] mb-1 font-normal text-text-primary opacity-90 tracking-[-0.01em]">Archive</h1>
        <p class="text-text-secondary opacity-70 text-sm">Completed tasks.</p>
      </div>
    </header>
    
    <div v-if="Object.keys(groupedTasks).length === 0" class="py-10 text-text-secondary text-left">
      <p>No completed tasks yet.</p>
    </div>

    <div class="flex flex-col">
      <div v-for="(tasks, dateStr) in groupedTasks" :key="dateStr" class="mb-10">
        <h3 class="text-[11px] font-semibold text-text-secondary opacity-60 uppercase tracking-widest mb-4 border-b border-border/50 pb-2">{{ dateStr }}</h3>
        
        <div class="flex flex-col">
          <div v-for="task in tasks" :key="task.id" class="flex items-start gap-4 py-4 border-b border-border/50 transition-all duration-200 first:border-t-0 px-2 rounded-md">
            <div class="shrink-0 mt-0.5">
              <Check class="text-success opacity-80" :size="18" :stroke-width="2"/>
            </div>
            
            <div class="flex-1 flex justify-between items-start gap-4">
              <span class="text-[15px] text-text-primary tracking-[-0.01em]">{{ task.title }}</span>
              <span class="text-[13px] text-text-secondary opacity-60 whitespace-nowrap text-right">
                {{ getDharmaName(task.dharma_id) }}
                <template v-if="task.project_id"> &middot; {{ getProjectName(task.project_id) }}</template>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
