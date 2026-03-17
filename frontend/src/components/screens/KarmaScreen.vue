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
  <div class="pt-[80px] flex flex-col mb-20 w-full">
    <!-- Header -->
    <header class="flex justify-between items-start mb-12">
      <div>
        <h1 class="font-serif text-[32px] mb-1 font-normal text-text-primary opacity-90 tracking-[-0.01em]">Karma</h1>
        <p class="text-text-secondary opacity-70 text-sm">Track your habits and behaviors.</p>
      </div>
      <div class="flex pt-2">
        <button class="flex items-center gap-1.5 bg-[#c68e3e] border-none px-3 py-1.5 rounded-md text-white text-[13px] font-medium hover:brightness-110 transition-all cursor-pointer" @click="emit('openNewKarmaModal')">
          <Plus :size="14" />
          <span>Log Karma</span>
        </button>
      </div>
    </header>
    
    <div v-if="Object.keys(groupedKarma).length === 0" class="py-10 text-text-secondary text-left">
      <p>No karma logged yet.</p>
    </div>

    <div class="flex flex-col">
      <div v-for="(entries, dateStr) in groupedKarma" :key="dateStr" class="mb-10">
        <h3 class="text-[11px] font-semibold text-text-secondary opacity-60 uppercase tracking-widest mb-4 border-b border-border/50 pb-2">{{ dateStr }}</h3>
        
        <div class="flex flex-col">
          <div v-for="entry in entries" :key="entry.id" class="flex items-center gap-4 py-4 border-b border-border/50 transition-all duration-200 first:border-t-0 hover:bg-surface/10 px-2 rounded-md">
            <div class="shrink-0">
              <PlusCircle v-if="entry.type === 'positive'" class="text-success opacity-80" :size="20" :stroke-width="1.5"/>
              <MinusCircle v-if="entry.type === 'negative'" class="text-danger opacity-80" :size="20" :stroke-width="1.5"/>
            </div>
            <div class="text-[15px] text-text-primary tracking-[-0.01em]">
              {{ entry.description }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
