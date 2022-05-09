type Issue = {
  id: number;
  title: string;
  url: string;
  commentCount: number;
  assigneesCount: number;
  gitHubCreatedAtFormatted: string;
  labels: string[];
};

export type { Issue };
