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
import { defaultParameters } from '../lib/default-values';
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
  const { locale, t } = useLocale();
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

  const [isDetailOpen, setIsDetailOpen] = useState(false);
  const [parameters, setParameters] = useState(defaultParameters);
  const [establishedParameters, setEstablishedParameters] = useState(defaultParameters);

  const fieldSpacing = 2.5;

  const detailOpenIcon = isDetailOpen ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />;

  const init = () => {
    fetchRepositories('init', parameters);
  };

  const search = () => {
    setEstablishedParameters(parameters);
    fetchRepositories('search', parameters);
  };

  const showmore = () => {
    fetchRepositories('showmore', establishedParameters);
  };

  useEffect(() => {
    fetchChoices();
    init();
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
        <meta property='og:title' content={t.HEADER_TITLE} />
        <meta property='og:description' content={t.HEADER_DESCRIPTION} />
        <meta property='og:type' content='website' />
        <meta property='og:url' content={`https://gitnavi.dev/${locale}`} />
        <meta property='og:image' content='https://gitnavi.dev/logo_large.png' />
        <meta property='og:site_name' content='gitnavi' />
      </Head>
      <NavBar />
      <Container sx={{ mb: 4 }} maxWidth='xl'>
        {/* <Box sx={{ textAlign: 'center' }}>
          <div>{JSON.stringify(parameters)}</div>
          <div>{JSON.stringify(establishedParameters)}</div>
        </Box> */}
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
                        values={parameters.languages}
                        onChange={(event, values) =>
                          setParameters({ ...parameters, languages: values })
                        }
                      />
                    ) : (
                      <MultipleChipsSelectField
                        options={languageChoices}
                        values={parameters.languages}
                        onChange={(event) =>
                          setParameters({
                            ...parameters,
                            languages: event.target.value as string[],
                          })
                        }
                        onChipDelete={(values) =>
                          setParameters({ ...parameters, languages: values })
                        }
                      />
                    )}
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.LABELS_FIELD_LABEL}</FieldLabel>
                    {useMediaQuery(theme.breakpoints.up('sm')) ? (
                      <MultipleChipsAutoCompleteField
                        options={labelChoices}
                        values={parameters.labels}
                        onChange={(event, values) =>
                          setParameters({ ...parameters, labels: values })
                        }
                      />
                    ) : (
                      <MultipleChipsSelectField
                        options={labelChoices}
                        values={parameters.labels}
                        onChange={(event) =>
                          setParameters({ ...parameters, labels: event.target.value as string[] })
                        }
                        onChipDelete={(values: string[]) =>
                          setParameters({ ...parameters, labels: values })
                        }
                      />
                    )}
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.ASSIGN_STATUS_FIELD_LABEL}</FieldLabel>
                    <AssignStatusField
                      value={parameters.assignStatus}
                      items={assignStatusChoices}
                      onChange={(event) => {
                        setParameters({ ...parameters, assignStatus: event.target.value });
                      }}
                    />
                  </Box>
                  <Box sx={{ mb: fieldSpacing }}>
                    <FieldLabel>{t.ORDER_BY_FIELD_LABEL}</FieldLabel>
                    <OrderByField
                      value={parameters.ordermetric}
                      items={ordermetricChoices}
                      onChange={(event) => {
                        setParameters({ ...parameters, ordermetric: event.target.value });
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
                        minValue={parameters.starCountLower}
                        maxValue={parameters.starCountUpper}
                        onChangeMin={(event) =>
                          setParameters({ ...parameters, starCountLower: event.target.value })
                        }
                        onChangeMax={(event) =>
                          setParameters({ ...parameters, starCountUpper: event.target.value })
                        }
                      />
                    </Box>
                    <Box sx={{ mb: fieldSpacing }}>
                      <FieldLabel>{t.FORK_COUNT_FIELD_LABEL}</FieldLabel>
                      <MinMaxNumberFields
                        minValue={parameters.forkCountLower}
                        maxValue={parameters.forkCountUpper}
                        onChangeMin={(event) =>
                          setParameters({ ...parameters, forkCountLower: event.target.value })
                        }
                        onChangeMax={(event) =>
                          setParameters({ ...parameters, forkCountUpper: event.target.value })
                        }
                      />
                    </Box>
                    <Box sx={{ mb: fieldSpacing }}>
                      <FieldLabel>{t.LICENSE_FIELD_LABEL}</FieldLabel>
                      <LicenseField
                        value={parameters.license}
                        items={licenseChoices}
                        onChange={(event) => {
                          setParameters({ ...parameters, license: event.target.value });
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
                        onClick={() => {
                          setParameters(defaultParameters);
                        }}
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
                        onClick={search}
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
                  onClick={showmore}
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
