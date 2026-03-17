<script setup lang="ts">
import { computed } from 'vue';
import { store } from '../../store';
import { Plus, Trash2 } from 'lucide-vue-next';

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
  <div class="screen">
    <div class="flex justify-between items-center mb-4">
      <h1>Dharmas</h1>
      <button class="btn btn-primary" @click="handleNewDharma"><Plus :size="16"/> New Dharma</button>
    </div>
    
    <div class="dharma-grid">
      <div v-for="dharma in dharmasWithDetails" :key="dharma.id" class="card dharma-card relative group">
        <div class="flex justify-between items-start">
          <h3>{{ dharma.name }}</h3>
          <button 
            @click="handleDeleteDharma(dharma.id, dharma.name)"
            class="text-red-400 hover:text-red-500 opacity-0 group-hover:opacity-100 transition-opacity bg-transparent border-none cursor-pointer p-1"
            title="Delete Dharma"
          >
            <Trash2 :size="16"/>
          </button>
        </div>
        
        <div class="stats mt-4">
          <div class="stat-item">
            <span class="stat-val">{{ dharma.projects.length }}</span>
            <span class="stat-label">Projects</span>
          </div>
          <div class="stat-item">
            <span class="stat-val">{{ dharma.directTasks.length }}</span>
            <span class="stat-label">Direct Tasks</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dharma-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
}
.dharma-card {
  display: flex;
  flex-direction: column;
}
.dharma-card h3 {
  color: var(--accent-color);
  margin-top: 0;
}
.stats {
  display: flex;
  gap: 16px;
  border-top: 1px solid var(--border-color);
  padding-top: 12px;
}
.stat-item {
  display: flex;
  flex-direction: column;
}
.stat-val {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
}
.stat-label {
  font-size: 12px;
  color: var(--text-secondary);
}
</style>
