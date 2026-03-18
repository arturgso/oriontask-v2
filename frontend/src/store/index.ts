import { reactive } from 'vue';
import { v4 as uuidv4 } from 'uuid';
import { CreateDharma, ListAllDharmas, DeleteDharma } from '../../wailsjs/go/dharmas/DharmaAppService';
import { CompleteTask, CreateTask, FindTasksByDharmaID, DeleteTask } from '../../wailsjs/go/tasks/TaskService';

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

// Initial Data
const dharmas: Dharma[] = [];
const projects: Project[] = [];
const milestones: Milestone[] = [];
const tasks: Task[] = [];
const karmaEntries: KarmaEntry[] = [];

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

      // Load tasks for each dharma
      this.tasks = [];
      for (const dharma of this.dharmas) {
        const dharmaTasks = await FindTasksByDharmaID(dharma.id as any);
        this.tasks.push(...dharmaTasks.map(t => ({
          id: t.id as unknown as string,
          dharma_id: t.dharmaId as unknown as string,
          title: t.title,
          description: t.description,
          status: (t.status === 'COMPLETED' || t.completedAt ? 'completed' : 'active') as 'active' | 'completed' | 'postponed',
          created_at: t.createdAt as unknown as string,
          completed_at: t.completedAt as unknown as string,
        })));
      }
    } catch (e) {
      console.error("Failed to load dharmas or tasks:", e);
    }
  },
  async addTask(taskData: Omit<Task, 'id' | 'created_at' | 'status'>) {
    try {
      const result = await CreateTask({
        title: taskData.title,
        description: taskData.description,
        dharmaId: taskData.dharma_id as any,
        status: 'TODO',
      } as any);

      this.tasks.push({
        id: result.id as unknown as string,
        dharma_id: result.dharmaId as unknown as string,
        title: result.title,
        description: result.description,
        status: 'active',
        created_at: result.createdAt as unknown as string,
      });
    } catch (e) {
      console.error("Failed to create task:", e);
    }
  },
  async completeTask(taskId: string) {
    try {
      const result = await CompleteTask(taskId as any);
      const task = this.tasks.find(t => t.id === taskId);
      if (task) {
        task.status = 'completed';
        task.completed_at = result.completedAt as unknown as string;
      }
    } catch (e) {
      console.error("Failed to complete task:", e);
    }
  },
  postponeTask(taskId: string) {
    // TODO: Call backend TaskService.PostponeTask() (if applicable)
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
      const newDharma = {
        id: result.id as unknown as string,
        name: result.name,
        created_at: result.created_at as unknown as string,
      };
      this.dharmas.push(newDharma);
    } catch (e) {
      console.error("Failed to create dharma:", e);
      throw e;
    }
  },
  async deleteDharma(id: string) {
    try {
      await DeleteDharma(id as any);
      this.dharmas = this.dharmas.filter(d => d.id !== id);
      this.tasks = this.tasks.filter(t => t.dharma_id !== id);
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

