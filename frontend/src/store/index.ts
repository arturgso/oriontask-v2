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

// Initial Mock Data (Except Dharmas)
const dharmas: Dharma[] = [];

const projects: Project[] = [
  { id: 'p1', dharma_id: 'd2', title: 'OrionTask', description: 'Productivity Application', created_at: new Date().toISOString() },
];

const milestones: Milestone[] = [
  { id: 'm1', project_id: 'p1', title: 'Simplification', description: 'Simplify application features', created_at: new Date().toISOString() },
];

const tasks: Task[] = [
  { id: 't1', dharma_id: 'd1', title: 'Clean the kitchen', description: '', status: 'active', created_at: new Date().toISOString() },
  { id: 't2', dharma_id: 'd2', project_id: 'p1', milestone_id: 'm1', title: 'Remove authentication module', description: '', status: 'active', created_at: new Date().toISOString() },
  { id: 't3', dharma_id: 'd1', title: 'Meditate for 10 minutes', description: '', status: 'active', created_at: new Date().toISOString() },
  { id: 't4', dharma_id: 'd2', project_id: 'p1', title: 'Design dark mode theme', description: '', status: 'active', created_at: new Date().toISOString() },
  { id: 't5', dharma_id: 'd2', title: 'Reply to emails', description: '', status: 'active', created_at: new Date().toISOString() },
  { id: 't6', dharma_id: 'd1', title: 'Read a book', description: '', status: 'active', created_at: new Date().toISOString() }, // Backlog
];

const karmaEntries: KarmaEntry[] = [
  { id: 'k1', type: 'positive', description: 'Meditated', created_at: new Date().toISOString() },
  { id: 'k2', type: 'negative', description: 'Procrastinated on project', created_at: new Date().toISOString() },
];

export const store = reactive({
  dharmas,
  projects,
  milestones,
  tasks,
  karmaEntries,

  // Actions
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
    this.tasks.push({
      ...taskData,
      id: uuidv4(),
      status: 'active',
      created_at: new Date().toISOString(),
    });
  },
  completeTask(taskId: string) {
    const task = this.tasks.find(t => t.id === taskId);
    if (task) {
      task.status = 'completed';
      task.completed_at = new Date().toISOString();
    }
  },
  postponeTask(taskId: string) {
    const task = this.tasks.find(t => t.id === taskId);
    if (task) {
      task.status = 'postponed';
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
