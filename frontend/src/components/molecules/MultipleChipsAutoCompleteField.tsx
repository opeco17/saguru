import { useLocale } from '../../hooks/locale';
import CheckBoxIcon from '@mui/icons-material/CheckBox';
import CheckBoxOutlineBlankIcon from '@mui/icons-material/CheckBoxOutlineBlank';
import Autocomplete from '@mui/material/Autocomplete';
import Checkbox from '@mui/material/Checkbox';
import Chip from '@mui/material/Chip';
import TextField from '@mui/material/TextField';

type MultipleChipsAutoCompleteProps = {
  options: string[];
  values: string[];
  onChange: any;
};

const icon = <CheckBoxOutlineBlankIcon fontSize='small' />;
const checkedIcon = <CheckBoxIcon fontSize='small' />;

const MultipleChipsAutoCompleteField = (props: MultipleChipsAutoCompleteProps) => {
  const { t } = useLocale();
  return (
    <>
      <Autocomplete
        multiple
        options={props.options}
        value={props.values}
        onChange={props.onChange}
        disableClearable
        renderTags={(values: readonly string[], getTagProps) =>
          values.map((value: string, index: number) => (
            <Chip
              variant='outlined'
              size='small'
              // @ts-ignore to use custom color
              color='greyChip'
              label={value === 'ALL' ? t.ALL : value}
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

export default MultipleChipsAutoCompleteField;
