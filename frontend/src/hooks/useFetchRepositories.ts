import client from '../lib/api-client';
import { Parameters } from '../types/parameters';
import { RepositoriesParams } from '../types/repositories-params';
import { RepositoriesResponseBody } from '../types/repositories-response-body';
import { Repository } from '../types/repository';
import { useState } from 'react';

const useFetchRepositories = () => {
  const [repositories, setRepositories] = useState([] as Repository[]);
  const [page, setPage] = useState(0);
  const [hasNext, setHasNext] = useState(false);
  const [isInitSearchLoading, setIsInitSearchLoading] = useState(false);
  const [isSearchLoading, setIsSearchLoading] = useState(false);
  const [isShowMoreLoading, setIsShowMoreLoading] = useState(false);
  const [fetchRepositoriesErrorMessage, setFetchRepositoriesErrorMessage] = useState('');

  const fetchRepositories = async (
    type: 'init' | 'search' | 'showmore',
    {
      languages,
      labels,
      assignStatus,
      ordermetric,
      license,
      starCountLower,
      starCountUpper,
      forkCountLower,
      forkCountUpper,
    }: Parameters,
  ) => {
    setFetchRepositoriesErrorMessage('');

    let targetPage: number;
    if (type === 'init') {
      targetPage = 0;
      setIsInitSearchLoading(true);
    } else if (type === 'search') {
      targetPage = 0;
      setIsSearchLoading(true);
    } else if (type === 'showmore') {
      targetPage = page;
      setIsShowMoreLoading(true);
    } else {
      throw 'Invalid type';
    }

    const params: RepositoriesParams = { page: targetPage };
    if (languages.length !== 0 && !languages.includes('ALL')) {
      params.languages = languages.join(',');
    }
    if (labels.length !== 0 && !labels.includes('ALL')) {
      params.labels = labels.join(',');
    }
    if (assignStatus !== 'ALL') {
      params.assigned = { ASSIGNED: true, UNASSIGNED: false }[assignStatus];
    }
    if (ordermetric) {
      params.orderby = ordermetric;
    }
    if (license !== 'ALL') {
      params.license = license;
    }
    if (starCountLower !== '') {
      params.star_count_lower = starCountLower;
    }
    if (starCountUpper !== '') {
      params.star_count_upper = starCountUpper;
    }
    if (forkCountLower !== '') {
      params.fork_count_lower = forkCountLower;
    }
    if (forkCountUpper !== '') {
      params.fork_count_upper = forkCountUpper;
    }

    try {
      const res = await client.get<RepositoriesResponseBody>('/repositories', { params: params });
      setHasNext(res.data.hasNext);
      if (type === 'init' || type === 'search') {
        setPage(1);
        setRepositories(res.data.items);
      } else if (type === 'showmore') {
        setPage((prev) => prev + 1);
        setRepositories((prev) => [...prev, ...res.data.items]);
      }
    } catch {
      setFetchRepositoriesErrorMessage('Failed to fetch repositories from API');
    } finally {
      if (type === 'init') {
        setIsInitSearchLoading(false);
      } else if (type === 'search') {
        setIsSearchLoading(false);
      } else if (type === 'showmore') {
        setIsShowMoreLoading(false);
      }
    }
  };
  return {
    repositories,
    hasNext,
    isInitSearchLoading,
    isSearchLoading,
    isShowMoreLoading,
    fetchRepositoriesErrorMessage,
    fetchRepositories,
  };
};

export default useFetchRepositories;
