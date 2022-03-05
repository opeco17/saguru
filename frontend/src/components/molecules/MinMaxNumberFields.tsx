import { useLocale } from '../../hooks/locale';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import { useState } from 'react';

type MinMaxNumberFieldsProps = {
  minValue: number | '';
  maxValue: number | '';
  onChangeMin: any;
  onChangeMax: any;
};

const MinMaxNumberFields = (props: MinMaxNumberFieldsProps) => {
  const { t } = useLocale();

  const [minInputError, setMinInputError] = useState(false);
  const [maxInputError, setMaxInputError] = useState(false);

  const [minInputErrorMessage, setMinInputErrorMessage] = useState('');
  const [maxInputErrorMessage, setMaxInputErrorMessage] = useState('');

  const onChangeMinHandler = (event: any) => {
    setMinInputError(false);
    setMinInputErrorMessage('');

    const value = event.target.value;
    if (value !== '' && value < 0) {
      setMinInputError(true);
      setMinInputErrorMessage(t.EQUAL_OR_GREATER_THAN_ZERO_ERROR_MESSAGE);
    }
    props.onChangeMin(event);
  };

  const onChangeMaxHandler = (event: any) => {
    setMaxInputError(false);
    setMaxInputErrorMessage('');

    const value = event.target.value;
    if (value !== '' && value < 0) {
      setMaxInputError(true);
      setMaxInputErrorMessage(t.EQUAL_OR_GREATER_THAN_ZERO_ERROR_MESSAGE);
    }
    props.onChangeMax(event);
  };

  return (
    <>
      <Grid container spacing={3}>
        <Grid item xs={6}>
          <TextField
            placeholder='Min'
            size='small'
            type='number'
            value={props.minValue}
            onChange={props.onChangeMin}
            error={minInputError}
            helperText={minInputErrorMessage}
          ></TextField>
        </Grid>
        <Grid item xs={6}>
          <TextField
            placeholder='Max'
            size='small'
            type='number'
            value={props.maxValue}
            onChange={props.onChangeMax}
            error={maxInputError}
            helperText={maxInputErrorMessage}
          ></TextField>
        </Grid>
      </Grid>
    </>
  );
};

export default MinMaxNumberFields;
