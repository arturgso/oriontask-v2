<script setup lang="ts">
import { computed, ref } from 'vue';
import { store } from '../../store';
import { PlusCircle, MinusCircle, Plus } from 'lucide-vue-next';

const groupedKarma = computed(() => {
  const groups: Record<string, typeof store.karmaEntries> = {};
  
  // Sort descending by date
  const sorted = [...store.karmaEntries].sort((a, b) => 
    new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
  );

  sorted.forEach(entry => {
    // group by simple date string
    const dateStr = new Date(entry.created_at).toLocaleDateString(undefined, {
      month: 'long', day: 'numeric', year: 'numeric'
    });
    if (!groups[dateStr]) groups[dateStr] = [];
    groups[dateStr].push(entry);
  });

  return groups;
});

const emit = defineEmits(['openNewKarmaModal']);
</script>

<template>
  <div class="screen">
    <div class="flex justify-between items-center mb-4">
      <h1>Karma</h1>
      <button class="btn btn-primary" @click="emit('openNewKarmaModal')">
        <Plus :size="16"/> Log Karma
      </button>
    </div>
    
    <div class="karma-log">
      <div v-for="(entries, dateStr) in groupedKarma" :key="dateStr" class="date-group">
        <h3 class="date-header">{{ dateStr }}</h3>
        
        <div class="entries">
          <div v-for="entry in entries" :key="entry.id" class="entry-item card flex items-center gap-4">
            <div class="entry-icon">
              <PlusCircle v-if="entry.type === 'positive'" class="text-positive" :size="20"/>
              <MinusCircle v-if="entry.type === 'negative'" class="text-negative" :size="20"/>
            </div>
            <div class="entry-desc">
              {{ entry.description }}
            </div>
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
}
.entries {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.entry-item {
  padding: 16px;
  margin-bottom: 0;
}
.text-positive { color: var(--success-color); }
.text-negative { color: var(--danger-color); }
.entry-desc {
  font-size: 15px;
}
</style>
