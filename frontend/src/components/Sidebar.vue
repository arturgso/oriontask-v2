<script setup lang="ts">
import { Circle, Folder, Compass, Sparkles, Box } from 'lucide-vue-next';

const props = defineProps<{ activeScreen: string }>();
const emit = defineEmits(['changeScreen']);

const navItems = [
  { id: 'now', label: 'Now', icon: Circle },
  { id: 'projects', label: 'Projects', icon: Folder },
  { id: 'dharmas', label: 'Dharmas', icon: Compass },
  { id: 'karma', label: 'Karma', icon: Sparkles },
  { id: 'archive', label: 'Archive', icon: Box },
];
</script>

<template>
  <aside class="w-[60px] bg-bg border-r border-border flex flex-col shrink-0 pt-5">
    <nav class="p-2 flex flex-col gap-3 items-center">
      <button 
        v-for="item in navItems" 
        :key="item.id"
        :class="[
          'flex items-center justify-center w-10 h-10 bg-transparent border-none rounded-sm cursor-pointer transition-all duration-200 hover:bg-surface-hover hover:text-text-primary', 
          activeScreen === item.id ? 'text-text-primary bg-surface' : 'text-text-secondary'
        ]"
        @click="emit('changeScreen', item.id)"
        :title="item.label"
      >
        <component :is="item.icon" :size="20" class="stroke-[1.5px]" />
      </button>
    </nav>
  </aside>
</template>
