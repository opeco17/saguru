import FieldLabel from '../components/atoms/FieldLabel';
import SubTitle from '../components/atoms/SubTitle';
import Title from '../components/atoms/Title';
import AssignStatusField from '../components/molecules/AssignStatusField';
import ErrorMessages from '../components/molecules/ErrorMessages';
import LicenseField from '../components/molecules/LicenseField';
import MinMaxNumberFields from '../components/molecules/MinMaxNumberFields';
import MultipleChipsAutoCompleteField from '../components/molecules/MultipleChipsAutoCompleteField';
import MultipleChipsSelectField from '../components/molecules/MultipleChipsSelectField';
import OrderByField from '../components/molecules/OrderByField';
import ScrollTopButton from '../components/molecules/ScrollTopButton';
import NavBar from '../components/organisms/NavBar';
import RepositoryCard from '../components/organisms/RepositoryCard';
import { useLocale } from '../hooks/locale';
import useFetchChoices from '../hooks/useFetchChoices';
import useFetchRepositories from '../hooks/useFetchRepositories';
import {
  defaultLabels,
  defaultLanguages,
  defaultLicense,
  defaultOrdermetric,
  defaultAssignStatus,
  defaultStarCountLower,
  defaultStarCountUpper,
  defaultForkCountLower,
  defaultForkCountUpper,
} from '../lib/default-values';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import LoadingButton from '@mui/lab/LoadingButton';
import { FormControl } from '@mui/material';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import Collapse from '@mui/material/Collapse';
import Container from '@mui/material/Container';
import Divider from '@mui/material/Divider';
import Fab from '@mui/material/Fab';
import Grid from '@mui/material/Grid';
import LinearProgress from '@mui/material/LinearProgress';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';
import Head from 'next/head';
import { useState, useEffect } from 'react';
import * as React from 'react';

const Index = () => {
  const theme = useTheme();
  const {
    labelChoices,
    languageChoices,
    licenseChoices,
    ordermetricChoices,
    assignStatusChoices,
    fetchChoicesErrorMessage,
    fetchChoices,
  } = useFetchChoices();

  const {
    repositories,
    hasNext,
    isInitSearchLoading,
    isSearchLoading,
    isShowMoreLoading,
    fetchRepositoriesErrorMessage,
    fetchRepositories,
  } = useFetchRepositories();

  const { t } = useLocale();
  const [isDetailOpen, setIsDetailOpen] = useState(false);

  const [languages, setLanguages] = useState(defaultLanguages);
  const [labels, setLabels] = useState(defaultLabels);
  const [assignStatus, setAssignStatus] = useState(defaultAssignStatus);
  const [ordermetric, setOrdermetric] = useState(defaultOrdermetric);
  const [license, setLicense] = useState(defaultLicense);
  const [starCountLower, setStarCountLower] = useState(defaultStarCountLower);
  const [starCountUpper, setStarCountUpper] = useState(defaultStarCountUpper);
  const [forkCountLower, setForkCountLower] = useState(defaultForkCountLower);
  const [forkCountUpper, setForkCountUpper] = useState(defaultForkCountUpper);

  const [establishedLanguages, setEstablishedLanguages] = useState(defaultLanguages);
  const [establishedLabels, setEstablishedLabels] = useState(defaultLabels);
  const [establishedAssignStatus, setEstablishedAssignStatus] = useState(defaultAssignStatus);
  const [establishedOrdermetric, setEstablishedOrdermetric] = useState(defaultOrdermetric);
  const [establishedLicense, setEstablishedLicense] = useState(defaultLicense);
  const [establishedStarCountLower, setEstablishedStarCountLower] = useState(defaultStarCountLower);
  const [establishedStarCountUpper, setEstablishedStarCountUpper] = useState(defaultStarCountUpper);
  const [establishedForkCountLower, setEstablishedForkCountLower] = useState(defaultForkCountLower);
  const [establishedForkCountUpper, setEstablishedForkCountUpper] = useState(defaultForkCountUpper);

  const fieldSpacing = 2.5;

  const detailOpenIcon = isDetailOpen ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />;

  const resetValues = () => {
    setLanguages(defaultLanguages);
    setLabels(defaultLabels);
    setAssignStatus(defaultAssignStatus);
    setOrdermetric(defaultOrdermetric);
    setLicense(defaultLicense);
    setStarCountLower(defaultStarCountLower);
    setStarCountUpper(defaultStarCountUpper);
    setForkCountLower(defaultForkCountLower);
    setForkCountUpper(defaultForkCountUpper);
  };

  const setParameters = () => {
    setEstablishedLanguages(languages);
    setEstablishedLabels(labels);
    setEstablishedAssignStatus(assignStatus);
    setEstablishedOrdermetric(ordermetric);
    setEstablishedLicense(license);
    setEstablishedStarCountLower(starCountLower);
    setEstablishedStarCountUpper(starCountUpper);
    setEstablishedForkCountLower(forkCountLower);
    setEstablishedForkCountUpper(forkCountUpper);
  };

  const fetchRepositoriesWrapper = (type: 'init' | 'search' | 'showmore') => {
    if (type === 'init') {
      setParameters();
    }
    return () =>
      fetchRepositories(
        type,
        establishedLanguages,
        establishedLabels,
        establishedAssignStatus,
        establishedOrdermetric,
        establishedLicense,
        establishedStarCountLower,
        establishedStarCountUpper,
        establishedForkCountLower,
        establishedForkCountUpper,
      );
  };

  useEffect(() => {
    fetchChoices();
  }, []);

  useEffect(() => {
    fetchRepositoriesWrapper('init')();
  }, []);

  let errorMessages: string[] = [];
  if (fetchChoicesErrorMessage) {
    errorMessages.push(fetchChoicesErrorMessage);
  }
  if (fetchRepositoriesErrorMessage) {
    errorMessages.push(fetchRepositoriesErrorMessage);
  }

  return (
    <>
      <Head>
        <title>{t.HEADER_TITLE} - gitnavi</title>
        <meta name='description' content={t.HEADER_DESCRIPTION}></meta>
      </Head>
      <NavBar />
      <Container sx={{ mb: 4 }}>
        <Box sx={{ mt: 4, textAlign: 'center', width: '100%' }} id='back-to-top-anchor'>
          <Title />
          <SubTitle />
        </Box>
        <Box>
          <Box sx={{ height: 25, mt: 1 }}>
            {/* @ts-ignore to use custom color */}
            {isInitSearchLoading || isSearchLoading || isShowMoreLoading ? (
              <LinearProgress color='info' />
            ) : null}
          </Box>
          <Box>
            {errorMessages.length === 0 ? null : <ErrorMessages errorMessages={errorMessages} />}
          </Box>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={5} md={4}>
              <Card variant='outlined' sx={{ px: 2, py: 2 }}>
                <FormControl fullWidth size='small'>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.LANGUAGES_FIELD_LABEL}</FieldLabel>
                    {useMediaQuery(theme.breakpoints.up('sm')) ? (
                      <MultipleChipsAutoCompleteField
                        options={languageChoices}
                        values={languages}
                        onChange={(event: any, values: string[]) => setLanguages(values)}
                      />
                    ) : (
                      <MultipleChipsSelectField
                        options={languageChoices}
                        values={languages}
                        onChange={(event: any) => setLanguages(event.target.value)}
                        onChipDelete={(values: string[]) => setLanguages(values)}
                      />
                    )}
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.LABELS_FIELD_LABEL}</FieldLabel>
                    {useMediaQuery(theme.breakpoints.up('sm')) ? (
                      <MultipleChipsAutoCompleteField
                        options={labelChoices}
                        values={labels}
                        onChange={(event: any, values: string[]) => setLabels(values)}
                      />
                    ) : (
                      <MultipleChipsSelectField
                        options={labelChoices}
                        values={labels}
                        onChange={(event: any) => setLabels(event.target.value)}
                        onChipDelete={(values: string[]) => setLabels(values)}
                      />
                    )}
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.ASSIGN_STATUS_FIELD_LABEL}</FieldLabel>
                    <AssignStatusField
                      value={assignStatus}
                      items={assignStatusChoices}
                      onChange={(event: any) => {
                        setAssignStatus(event.target.value);
                      }}
                    />
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.ORDER_BY_FIELD_LABEL}</FieldLabel>
                    <OrderByField
                      value={ordermetric}
                      items={ordermetricChoices}
                      onChange={(event: any) => {
                        setOrdermetric(event.target.value);
                      }}
                    />
                  </Box>
                  <Box>
                    <Button
                      variant='text'
                      size='small'
                      // @ts-ignore to use custom color
                      color='greyChip'
                      sx={{ pl: 1, py: 0.1, mb: 0.5, width: '100px', float: 'right' }}
                      onClick={() => setIsDetailOpen((prev) => !prev)}
                    >
                      {t.DETAIL_BUTTON_LABEL} {detailOpenIcon}
                    </Button>
                  </Box>
                  <Divider sx={{ mb: 2 }} />
                  <Collapse orientation='vertical' in={isDetailOpen}>
                    <Box sx={{ mb: fieldSpacing }}>
                      <FieldLabel>{t.STAR_COUNT_FIELD_LABEL}</FieldLabel>
                      <MinMaxNumberFields
                        minValue={starCountLower}
                        maxValue={starCountUpper}
                        onChangeMin={(event: any) => setStarCountLower(event.target.value)}
                        onChangeMax={(event: any) => setStarCountUpper(event.target.value)}
                      />
                    </Box>
                    <Box sx={{ mb: fieldSpacing }}>
                      <FieldLabel>{t.FORK_COUNT_FIELD_LABEL}</FieldLabel>
                      <MinMaxNumberFields
                        minValue={forkCountLower}
                        maxValue={forkCountUpper}
                        onChangeMin={(event: any) => setForkCountLower(event.target.value)}
                        onChangeMax={(event: any) => setForkCountUpper(event.target.value)}
                      />
                    </Box>
                    <Box sx={{ mb: fieldSpacing }}>
                      <FieldLabel>{t.LICENSE_FIELD_LABEL}</FieldLabel>
                      <LicenseField
                        value={license}
                        items={licenseChoices}
                        onChange={(event: any) => {
                          setLicense(event.target.value);
                        }}
                      />
                    </Box>
                  </Collapse>
                  <Grid container spacing={2} sx={{ flexGrow: 1 }}>
                    <Grid item xs={6} sx={{ textAlign: 'right' }}>
                      <Button
                        variant='outlined'
                        color='info'
                        disableElevation
                        onClick={resetValues}
                        sx={{ py: 0.7, width: '100px' }}
                      >
                        {t.RESET_BUTTON_LABEL}
                      </Button>
                    </Grid>
                    <Grid item xs={6} sx={{ textAlign: 'left' }}>
                      <LoadingButton
                        loading={isSearchLoading}
                        variant='contained'
                        color='info'
                        disableElevation
                        sx={{ py: 0.7, width: '100px' }}
                        onClick={fetchRepositoriesWrapper('search')}
                      >
                        {t.SEARCH_BUTTON_LABEL}
                      </LoadingButton>
                    </Grid>
                  </Grid>
                </FormControl>
              </Card>
            </Grid>
            <Grid item xs={12} sm={7} md={8}>
              {repositories.map((repository) => {
                return <RepositoryCard repository={repository} key={repository.id} />;
              })}
              <Box sx={{ textAlign: 'center' }}>
                <LoadingButton
                  loading={isShowMoreLoading}
                  variant='contained'
                  startIcon={<KeyboardArrowDownIcon />}
                  color='info'
                  disabled={!hasNext}
                  disableElevation
                  sx={{ px: 2.5, py: 0.7 }}
                  onClick={fetchRepositoriesWrapper('showmore')}
                >
                  {t.SHOW_MORE_BUTTON_LABEL}
                </LoadingButton>
              </Box>
            </Grid>
          </Grid>
        </Box>
      </Container>
      <ScrollTopButton>
        {/* @ts-ignore to use custom color */}
        <Fab color='info' sx={{ mr: 1 }}>
          <KeyboardArrowUpIcon />
        </Fab>
      </ScrollTopButton>
    </>
  );
};

export default Index;
