<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { ChevronDown, Check } from 'lucide-vue-next';

interface Option {
  id: string;
  name: string;
}

const props = defineProps<{
  modelValue: string;
  options: Option[];
  placeholder: string;
  label?: string;
}>();

const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);
const selectRef = ref<HTMLElement | null>(null);

const toggle = () => {
  isOpen.value = !isOpen.value;
};

const selectOption = (id: string) => {
  emit('update:modelValue', id);
  isOpen.value = false;
};

const selectedOptionName = computed(() => {
  const option = props.options.find(o => o.id === props.modelValue);
  return option ? option.name : '';
});

const handleClickOutside = (event: MouseEvent) => {
  if (selectRef.value && !selectRef.value.contains(event.target as Node)) {
    isOpen.value = false;
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<template>
  <div class="custom-select-container w-full" ref="selectRef">
    <label v-if="label" class="block mb-1.5 text-[13px] text-text-secondary">{{ label }}</label>
    
    <div class="relative">
      <!-- Trigger Button -->
      <button 
        type="button"
        @click="toggle"
        class="w-full bg-surface border border-border text-text-primary px-3 py-[10px] rounded-sm font-sans text-[14px] text-left flex justify-between items-center transition-all duration-200 hover:border-border/80 focus:border-accent outline-none cursor-pointer"
        :class="{ 'border-accent ring-1 ring-accent/20': isOpen }"
      >
        <span :class="{ 'opacity-40': !modelValue }">
          {{ selectedOptionName || placeholder }}
        </span>
        <ChevronDown 
          :size="16" 
          class="transition-transform duration-200 opacity-50"
          :class="{ 'rotate-180': isOpen }"
        />
      </button>

      <!-- Options Dropdown -->
      <div 
        v-if="isOpen"
        class="absolute z-[110] w-full mt-1 bg-surface border border-border rounded-md shadow-xl overflow-hidden"
      >
        <div class="max-h-[240px] overflow-y-auto py-1 custom-scrollbar">
          <button
            v-for="option in options"
            :key="option.id"
            type="button"
            @click="selectOption(option.id)"
            class="w-full text-left px-3 py-2.5 text-[14px] transition-colors flex items-center justify-between group cursor-pointer border-none bg-transparent"
            :class="modelValue === option.id ? 'bg-accent/10 text-accent' : 'text-text-primary hover:bg-surface-hover'"
          >
            <span>{{ option.name }}</span>
            <Check v-if="modelValue === option.id" :size="14" />
          </button>
          
          <div v-if="options.length === 0" class="px-3 py-4 text-center text-text-secondary opacity-50 text-[13px]">
            No options available
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--border-color);
  border-radius: 10px;
}
</style>
