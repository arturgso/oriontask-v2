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
  <aside class="w-[60px] md:w-[200px] bg-bg border-r border-border flex flex-col shrink-0 pt-5 transition-all duration-300">
    <nav class="px-2 flex flex-col gap-1.5 items-center md:items-stretch">
      <button 
        v-for="item in navItems" 
        :key="item.id"
        :class="[
          'flex items-center gap-3 w-10 md:w-full h-10 md:px-3 bg-transparent border-none rounded-md cursor-pointer transition-all duration-200 hover:bg-surface-hover hover:text-text-primary justify-center md:justify-start', 
          activeScreen === item.id ? 'text-text-primary bg-surface' : 'text-text-secondary'
        ]"
        @click="emit('changeScreen', item.id)"
        :title="item.label"
      >
        <component :is="item.icon" :size="18" class="stroke-[2px] shrink-0" />
        <span class="hidden md:inline text-[14px] font-medium tracking-wide">{{ item.label }}</span>
      </button>
    </nav>
  </aside>
</template>
