<script setup lang="ts">
import { computed } from 'vue';
import { store } from '../../store';
import { Plus, Trash2, Folder, ListTodo, ChevronRight } from 'lucide-vue-next';

const emit = defineEmits(['openNewDharmaModal', 'openConfirm']);

const dharmasWithDetails = computed(() => {
  return store.dharmas.map(d => {
    return {
      ...d,
      projects: store.projects.filter(p => p.dharma_id === d.id),
      directTasks: store.tasks.filter(t => t.dharma_id === d.id && !t.project_id && t.status === 'active')
    }
  });
});

const handleNewDharma = () => {
  emit('openNewDharmaModal');
};

const handleDeleteDharma = (id: string, name: string) => {
  emit('openConfirm', {
    title: 'Delete Dharma',
    message: `Are you sure you want to delete "${name}"? This action cannot be undone.`,
    confirmLabel: 'Delete',
    isDestructive: true,
    onConfirm: async () => {
      await store.deleteDharma(id);
    }
  });
};
</script>

<template>
  <div class="pt-[80px] flex flex-col mb-20 w-full">
    <!-- Header -->
    <header class="flex justify-between items-start mb-12">
      <div>
        <h1 class="font-serif text-[32px] mb-1 font-normal text-text-primary opacity-90 tracking-[-0.01em]">Dharmas</h1>
        <p class="text-text-secondary opacity-70 text-sm">Your core areas of focus and responsibility.</p>
      </div>
      <div class="flex pt-2">
        <button class="flex items-center gap-1.5 bg-[#c68e3e] border-none px-3 py-1.5 rounded-md text-white text-[13px] font-medium hover:brightness-110 transition-all cursor-pointer" @click="handleNewDharma">
          <Plus :size="14" />
          <span>New Dharma</span>
        </button>
      </div>
    </header>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div v-for="dharma in dharmasWithDetails" :key="dharma.id" 
        class="group bg-surface/20 border border-border/40 rounded-2xl p-7 flex flex-col transition-all duration-300 hover:bg-surface/40 hover:border-border/80 hover:-translate-y-1 hover:shadow-xl hover:shadow-black/10 cursor-default"
      >
        <!-- Card Header -->
        <div class="flex justify-between items-start mb-8">
          <div class="flex flex-col gap-1">
            <h3 class="text-[18px] font-medium text-text-primary tracking-tight leading-tight group-hover:text-accent transition-colors">
              {{ dharma.name }}
            </h3>
            <span class="text-[12px] text-text-secondary opacity-50 font-medium uppercase tracking-widest">
              Core Dharma
            </span>
          </div>
          
          <button 
            @click.stop="handleDeleteDharma(dharma.id, dharma.name)"
            class="text-text-secondary/30 hover:text-danger opacity-0 group-hover:opacity-100 transition-all bg-transparent border-none cursor-pointer p-2 rounded-full hover:bg-danger/10 -mt-2 -mr-2"
            title="Delete Dharma"
          >
            <Trash2 :size="16" :stroke-width="1.5"/>
          </button>
        </div>
        
        <!-- Stats Section -->
        <div class="flex items-center gap-8 mt-auto pt-6 border-t border-border/30">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-lg bg-surface flex items-center justify-center text-text-secondary opacity-60 group-hover:opacity-100 transition-opacity">
              <Folder :size="14" :stroke-width="2" />
            </div>
            <div class="flex flex-col">
              <span class="text-[15px] font-semibold text-text-primary leading-none mb-1">{{ dharma.projects.length }}</span>
              <span class="text-[11px] text-text-secondary opacity-60 font-medium">Projects</span>
            </div>
          </div>

          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-lg bg-surface flex items-center justify-center text-text-secondary opacity-60 group-hover:opacity-100 transition-opacity">
              <ListTodo :size="14" :stroke-width="2" />
            </div>
            <div class="flex flex-col">
              <span class="text-[15px] font-semibold text-text-primary leading-none mb-1">{{ dharma.directTasks.length }}</span>
              <span class="text-[11px] text-text-secondary opacity-60 font-medium">Tasks</span>
            </div>
          </div>

          <!-- Explore Hint -->
          <div class="ml-auto opacity-0 group-hover:opacity-40 transition-opacity">
            <ChevronRight :size="18" class="text-text-primary" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
