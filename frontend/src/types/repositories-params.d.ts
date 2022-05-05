type RepositoriesParams = {
  page?: number;
  labels?: string;
  isAssigned?: boolean;
  languages?: string;
  starCountLower?: number;
  starCountUpper?: number;
  forkCountLower?: number;
  forkCountUpper?: number;
  license?: string;
  orderby?: string;
  keyword?: string;
};

export type { RepositoriesParams };
