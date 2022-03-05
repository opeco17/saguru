import client from '../lib/api-client';
import {
  defaultLabels,
  defaultLanguages,
  defaultLicense,
  defaultOrdermetric,
} from '../lib/default-values';
import {
  LabelsResponseBody,
  LanguagesResponseBody,
  LicensesResponseBody,
  OrdermetricsResponseBody,
} from '../types/choices';
import { useState } from 'react';

const useFetchChoices = () => {
  const [labelChoices, setLabelChoices] = useState([...defaultLabels] as string[]);
  const [languageChoices, setLanguageChoices] = useState([...defaultLanguages] as string[]);
  const [licenseChoices, setLicenseChoices] = useState([defaultLicense] as string[]);
  const [ordermetricChoices, setOrdermetricChoices] = useState([defaultOrdermetric] as string[]);
  const assignStatusChoices = ['ALL', 'ASSIGNED', 'UNASSIGNED'];
  const [fetchChoicesErrorMessage, setFetchChoicesErrorMessage] = useState('');

  const fetchChoices = async () => {
    setFetchChoicesErrorMessage('');

    try {
      const results = [];
      results.push(
        client.get<LabelsResponseBody>('/labels').then((res) => {
          setLabelChoices(['ALL', ...res.data.items]);
        }),
      );
      results.push(
        client.get<LanguagesResponseBody>('/languages').then((res) => {
          setLanguageChoices(['ALL', ...res.data.items]);
        }),
      );
      results.push(
        client.get<LicensesResponseBody>('/licenses').then((res) => {
          setLicenseChoices(['ALL', ...res.data.items]);
        }),
      );
      results.push(
        client.get<OrdermetricsResponseBody>('/ordermetrics').then((res) => {
          setOrdermetricChoices(res.data.items);
        }),
      );
      await Promise.all(results);
    } catch {
      setFetchChoicesErrorMessage('Failed to fetch choices from API');
    }
  };
  return {
    labelChoices,
    languageChoices,
    licenseChoices,
    ordermetricChoices,
    assignStatusChoices,
    fetchChoicesErrorMessage,
    fetchChoices,
  };
};

export default useFetchChoices;
