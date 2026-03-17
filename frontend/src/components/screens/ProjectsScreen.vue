<script setup lang="ts">
import { computed } from 'vue';
import { store } from '../../store';
import { Plus, ChevronRight } from 'lucide-vue-next';

const dharmasWithProjects = computed(() => {
  return store.dharmas.map(d => {
    const dharmaProjects = store.projects.filter(p => p.dharma_id === d.id).map(p => {
      const pMilestones = store.milestones.filter(m => m.project_id === p.id).map(m => {
        const mTasks = store.tasks.filter(t => t.milestone_id === m.id && t.status === 'active');
        return {
          ...m,
          taskCount: mTasks.length
        }
      });
      
      const directTasks = store.tasks.filter(t => t.project_id === p.id && !t.milestone_id && t.status === 'active');
      const totalActiveTasks = directTasks.length + pMilestones.reduce((acc, m) => acc + m.taskCount, 0);
      
      return {
        ...p,
        milestones: pMilestones,
        totalActiveTasks
      }
    });

    return {
      ...d,
      projects: dharmaProjects
    }
  }).filter(d => d.projects.length > 0);
});
</script>

<template>
  <div class="pt-[80px] flex flex-col mb-20">
    <!-- Header -->
    <header class="flex justify-between items-start mb-12">
      <div>
        <h1 class="font-serif text-[32px] mb-1 font-normal text-text-primary opacity-90 tracking-[-0.01em]">Projects</h1>
        <p class="text-text-secondary opacity-70 text-sm">Organized by dharma.</p>
      </div>
      <div class="flex gap-3 pt-2">
        <button class="flex items-center gap-1.5 bg-bg border border-border px-3 py-1.5 rounded-md text-text-primary text-[13px] font-medium hover:bg-surface-hover transition-colors">
          <Plus :size="14" />
          <span>Milestone</span>
        </button>
        <button class="flex items-center gap-1.5 bg-[#c68e3e] border-none px-3 py-1.5 rounded-md text-white text-[13px] font-medium hover:brightness-110 transition-all">
          <Plus :size="14" />
          <span>Project</span>
        </button>
      </div>
    </header>
    
    <!-- Dharma Groups -->
    <div v-for="dharma in dharmasWithProjects" :key="dharma.id" class="mb-10">
      <h2 class="text-[11px] font-semibold text-text-secondary opacity-60 uppercase tracking-widest mb-4">
        {{ dharma.name }}
      </h2>
      
      <div v-for="project in dharma.projects" :key="project.id" class="bg-surface/30 border border-border/50 rounded-xl p-6 mb-6">
        <h3 class="text-base font-semibold text-text-primary mb-4">{{ project.title }}</h3>
        
        <!-- Milestones list -->
        <div class="flex flex-col gap-2.5 mb-5">
          <div v-for="milestone in project.milestones" :key="milestone.id" class="flex items-center gap-2 group cursor-pointer">
            <ChevronRight :size="14" class="text-text-secondary opacity-40 group-hover:opacity-100 transition-opacity" />
            <span class="text-[14px] text-text-secondary group-hover:text-text-primary transition-colors">{{ milestone.title }}</span>
            <span v-if="milestone.taskCount > 0" class="text-[10px] font-bold bg-border/40 text-text-secondary px-1.5 py-0.5 rounded-full min-w-[18px] text-center">
              {{ milestone.taskCount }}
            </span>
          </div>
        </div>

        <!-- Meta and Actions -->
        <div class="flex flex-col gap-2">
          <span class="text-[13px] text-text-secondary opacity-60">
            {{ project.totalActiveTasks }} active task{{ project.totalActiveTasks !== 1 ? 's' : '' }}
          </span>
          <button class="flex items-center gap-1 text-[#c68e3e] text-[13px] font-medium bg-transparent border-none p-0 w-fit hover:brightness-125 transition-all">
            <Plus :size="14" />
            <span>Add task</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
