import { useLocale } from '../../hooks/locale';
import CheckBoxIcon from '@mui/icons-material/CheckBox';
import CheckBoxOutlineBlankIcon from '@mui/icons-material/CheckBoxOutlineBlank';
import Autocomplete from '@mui/material/Autocomplete';
import Checkbox from '@mui/material/Checkbox';
import Chip from '@mui/material/Chip';
import TextField from '@mui/material/TextField';

type MultipleChipsFieldProps = {
  options: string[];
  value: string[];
  onChange: any;
};

const icon = <CheckBoxOutlineBlankIcon fontSize='small' />;
const checkedIcon = <CheckBoxIcon fontSize='small' />;

const MultipleChipsField = (props: MultipleChipsFieldProps) => {
  const { t } = useLocale();
  return (
    <>
      <Autocomplete
        multiple
        id='tags-outlined'
        options={props.options}
        value={props.value}
        onChange={props.onChange}
        renderTags={(value: readonly string[], getTagProps) =>
          value.map((option: string, index: number) => (
            <Chip
              variant='outlined'
              size='small'
              // @ts-ignore to use custom color
              color='greyChip'
              label={option === 'ALL' ? t.ALL : option}
              {...getTagProps({ index })}
              key={index}
              sx={{ px: 0.5 }}
            />
          ))
        }
        renderOption={(props, option, { selected }) => (
          <li {...props}>
            <Checkbox
              icon={icon}
              checkedIcon={checkedIcon}
              style={{ marginRight: 8 }}
              checked={selected}
            />
            {option === 'ALL' ? t.ALL : option}
          </li>
        )}
        renderInput={(params) => <TextField {...params} />}
      />
    </>
  );
};

export default MultipleChipsField;
