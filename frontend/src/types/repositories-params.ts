type RepositoriesParams = {
  page?: number;
  labels?: string;
  assigned?: boolean;
  languages?: string;
  star_count_lower?: number;
  star_count_upper?: number;
  fork_count_lower?: number;
  fork_count_upper?: number;
  license?: string;
  orderby?: string;
};

export type { RepositoriesParams };
