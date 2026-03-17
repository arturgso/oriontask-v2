<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { store } from './store';
import Sidebar from './components/Sidebar.vue';
import NowScreen from './components/screens/NowScreen.vue';
import ProjectsScreen from './components/screens/ProjectsScreen.vue';
import DharmasScreen from './components/screens/DharmasScreen.vue';
import KarmaScreen from './components/screens/KarmaScreen.vue';
import ArchiveScreen from './components/screens/ArchiveScreen.vue';
import NewTaskModal from './components/modals/NewTaskModal.vue';
import NewKarmaModal from './components/modals/NewKarmaModal.vue';
import NewDharmaModal from './components/modals/NewDharmaModal.vue';
import ConfirmModal from './components/modals/ConfirmModal.vue';
import { Plus } from 'lucide-vue-next';

onMounted(() => {
  store.loadDharmas();
});

const currentScreen = ref('now');
const isTaskModalOpen = ref(false);
const isKarmaModalOpen = ref(false);
const isDharmaModalOpen = ref(false);

// Confirm Modal State
const isConfirmModalOpen = ref(false);
const confirmConfig = ref({
  title: '',
  message: '',
  confirmLabel: '',
  isDestructive: false,
  onConfirm: () => {}
});

const openConfirm = (config: any) => {
  confirmConfig.value = { ...config };
  isConfirmModalOpen.value = true;
};

const handleConfirm = async () => {
  await confirmConfig.value.onConfirm();
  isConfirmModalOpen.value = false;
};

const screens: Record<string, any> = {
  now: NowScreen,
  projects: ProjectsScreen,
  dharmas: DharmasScreen,
  karma: KarmaScreen,
  archive: ArchiveScreen
};

const currentComponent = computed(() => screens[currentScreen.value] || NowScreen);
</script>

<template>
  <div class="flex h-full bg-bg text-text-primary w-full overflow-hidden">
    <Sidebar :activeScreen="currentScreen" @changeScreen="screen => currentScreen = screen" />
    
    <main class="flex-1 h-full overflow-y-auto bg-bg flex justify-center">
      <div class="w-full max-w-[700px] flex flex-col px-6">
        <component 
          :is="currentComponent" 
          @openNewKarmaModal="isKarmaModalOpen = true"
          @openNewTaskModal="isTaskModalOpen = true"
          @openNewDharmaModal="isDharmaModalOpen = true"
          @openConfirm="openConfirm"
        />
      </div>
    </main>

    <!-- Floating Action Button -->
    <button 
      class="fixed bottom-8 right-8 bg-accent text-white border-none rounded-full py-4 px-6 flex items-center gap-3 cursor-pointer shadow-[0_4px_12px_rgba(0,0,0,0.3)] transition-all duration-200 z-50 font-sans text-[15px] font-medium hover:-translate-y-[2px] hover:bg-accent-hover"
      @click="isTaskModalOpen = true" 
      title="New Task"
    >
      <Plus :size="24" />
      <span class="inline">New Task</span>
    </button>

    <!-- Modals -->
    <NewTaskModal v-if="isTaskModalOpen" @close="isTaskModalOpen = false" />
    <NewKarmaModal v-if="isKarmaModalOpen" @close="isKarmaModalOpen = false" />
    <NewDharmaModal v-if="isDharmaModalOpen" @close="isDharmaModalOpen = false" />
    <ConfirmModal 
      v-if="isConfirmModalOpen" 
      :title="confirmConfig.title"
      :message="confirmConfig.message"
      :confirmLabel="confirmConfig.confirmLabel"
      :isDestructive="confirmConfig.isDestructive"
      @confirm="handleConfirm"
      @close="isConfirmModalOpen = false" 
    />
  </div>
</template>
