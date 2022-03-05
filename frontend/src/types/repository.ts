import { Issue } from './issue';

type Repository = {
  id: number;
  name: string;
  url: string;
  description: string;
  starCount: number;
  forkCount: number;
  openIssueCount: number;
  topics: string;
  license: string;
  language: string;
  issues: Issue[];
};

export type { Repository };
