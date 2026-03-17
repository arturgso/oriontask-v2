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
  <div class="screen">
    <h1>Archive</h1>
    
    <div v-if="Object.keys(groupedTasks).length === 0" class="empty-state">
      <p>No completed tasks yet.</p>
    </div>

    <div class="archive-list">
      <div v-for="(tasks, dateStr) in groupedTasks" :key="dateStr" class="date-group">
        <h3 class="date-header">{{ dateStr }}</h3>
        
        <div class="entries">
          <div v-for="task in tasks" :key="task.id" class="task-item flex items-center gap-3">
            <Check class="text-success" :size="16"/>
            <span class="task-title">{{ task.title }}</span>
            <span class="meta text-muted">
              {{ getDharmaName(task.dharma_id) }}
              <template v-if="task.project_id"> &middot; {{ getProjectName(task.project_id) }}</template>
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.screen {
  padding: 40px;
  max-width: 800px;
  margin: 0 auto;
}
.date-group {
  margin-bottom: 32px;
}
.date-header {
  font-size: 14px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-bottom: 12px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 8px;
}
.entries {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.task-item {
  padding: 8px 0;
}
.text-success { color: var(--success-color); }
.task-title {
  font-size: 15px;
  color: var(--text-primary);
}
.meta {
  font-size: 13px;
  margin-left: auto;
}
.text-muted { color: var(--text-secondary); }
.empty-state { text-align: center; color: var(--text-secondary); margin-top: 40px; }
</style>
