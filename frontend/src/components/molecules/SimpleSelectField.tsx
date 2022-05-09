import { useLocale } from '../../hooks/locale';
import InputText from '../atoms/InputTest';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import { SelectChangeEvent } from '@mui/material';
import MenuItem from '@mui/material/MenuItem';

type SimpleSelectFieldProps = {
  value: string;
  items: string[];
  onChange: (event: SelectChangeEvent) => void;
};

const SimpleSelectField = ({ value, items, onChange }: SimpleSelectFieldProps) => {
  const { t } = useLocale();

  return (
    <SimpleSelectWrapper value={value} onChange={onChange}>
      {items.map((each) => {
        return (
          <MenuItem value={each} key={each}>
            <InputText>{t[each]}</InputText>
          </MenuItem>
        );
      })}
    </SimpleSelectWrapper>
  );
};

export default SimpleSelectField;
