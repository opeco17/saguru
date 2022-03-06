import { useLocale } from '../../hooks/locale';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import { SelectChangeEvent } from '@mui/material';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';

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
            <Typography>{t[each]}</Typography>
          </MenuItem>
        );
      })}
    </SimpleSelectWrapper>
  );
};

export default SimpleSelectField;
