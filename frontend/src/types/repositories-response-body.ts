import { Repository } from './repository';

type RepositoriesResponseBody = {
  items: Repository[];
  hasNext: boolean;
};

export type { RepositoriesResponseBody };
