import { reactive } from 'vue';
import { v4 as uuidv4 } from 'uuid';
import { CreateDharma, ListAllDharmas, DeleteDharma } from '../../wailsjs/go/dharmas/DharmaAppService';

export interface Dharma {
  id: string;
  name: string;
  created_at: string;
}

export interface Project {
  id: string;
  dharma_id: string;
  title: string;
  description: string;
  created_at: string;
}

export interface Milestone {
  id: string;
  project_id: string;
  title: string;
  description: string;
  created_at: string;
}

export interface Task {
  id: string;
  dharma_id: string;
  project_id?: string;
  milestone_id?: string;
  title: string;
  description: string;
  status: 'active' | 'completed' | 'postponed';
  created_at: string;
  completed_at?: string;
}

export interface KarmaEntry {
  id: string;
  type: 'positive' | 'negative';
  description: string;
  created_at: string;
}

// Initial Data (Loaded from storage or empty)
const dharmas: Dharma[] = [];
const projects: Project[] = [];
const milestones: Milestone[] = [];

const LOCAL_STORAGE_TASKS_KEY = 'oriontask_tasks';

const loadTasks = (): Task[] => {
  const stored = localStorage.getItem(LOCAL_STORAGE_TASKS_KEY);
  if (stored) {
    try {
      return JSON.parse(stored);
    } catch (e) {
      console.error('Failed to parse tasks from localStorage', e);
    }
  }
  return [];
};

const tasks: Task[] = loadTasks();

const karmaEntries: KarmaEntry[] = [];

export const store = reactive({
  dharmas,
  projects,
  milestones,
  tasks,
  karmaEntries,

  // Actions
  saveTasks() {
    localStorage.setItem(LOCAL_STORAGE_TASKS_KEY, JSON.stringify(this.tasks));
  },
  async loadDharmas() {
    try {
      const data = await ListAllDharmas();
      this.dharmas = data.map(d => ({
        id: d.id as unknown as string,
        name: d.name,
        created_at: d.created_at as unknown as string,
      }));
    } catch (e) {
      console.error("Failed to load dharmas:", e);
    }
  },
  addTask(taskData: Omit<Task, 'id' | 'created_at' | 'status'>) {
    // TODO: Call backend TaskService.CreateTask()
    this.tasks.push({
      ...taskData,
      id: uuidv4(),
      status: 'active',
      created_at: new Date().toISOString(),
    });
    this.saveTasks();
  },
  completeTask(taskId: string) {
    // TODO: Call backend TaskService.CompleteTask()
    const task = this.tasks.find(t => t.id === taskId);
    if (task) {
      task.status = 'completed';
      task.completed_at = new Date().toISOString();
      this.saveTasks();
    }
  },
  postponeTask(taskId: string) {
    // TODO: Call backend TaskService.PostponeTask() (if applicable)
    const task = this.tasks.find(t => t.id === taskId);
    if (task) {
      task.status = 'postponed';
      this.saveTasks();
    }
  },
  addKarma(entry: Omit<KarmaEntry, 'id' | 'created_at'>, dateStr: string) {
    this.karmaEntries.push({
      ...entry,
      id: uuidv4(),
      created_at: dateStr || new Date().toISOString(),
    });
  },
  async addDharma(name: string) {
    try {
      const result = await CreateDharma(name);
      this.dharmas.push({
        id: result.id as unknown as string,
        name: result.name,
        created_at: result.created_at as unknown as string,
      });
    } catch (e) {
      console.error("Failed to create dharma:", e);
      throw e;
    }
  },
  async deleteDharma(id: string) {
    try {
      // The API expects a uuid string mapped to number array or simply the string depending on bindings
      await DeleteDharma(id as any);
      this.dharmas = this.dharmas.filter(d => d.id !== id);
    } catch (e) {
      console.error("Failed to delete dharma:", e);
      throw e;
    }
  },
  addProject(dharma_id: string, title: string, description: string = '') {
    this.projects.push({
      id: uuidv4(),
      dharma_id,
      title,
      description,
      created_at: new Date().toISOString(),
    });
  },
  addMilestone(project_id: string, title: string, description: string = '') {
    this.milestones.push({
      id: uuidv4(),
      project_id,
      title,
      description,
      created_at: new Date().toISOString(),
    });
  }
});
