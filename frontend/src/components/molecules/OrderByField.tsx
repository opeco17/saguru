import { useLocale } from '../../hooks/locale';
import InputText from '../atoms/InputTest';
import SimpleSelectWrapper from '../atoms/SimpleSelectWrapper';
import { SelectChangeEvent } from '@mui/material';
import MenuItem from '@mui/material/MenuItem';

type OrderByFieldProps = {
  value: string;
  items: string[];
  onChange: (event: SelectChangeEvent<string>) => void;
};

const OrderByField = ({ value, items, onChange }: OrderByFieldProps) => {
  const { t } = useLocale();

  return (
    <>
      <SimpleSelectWrapper value={value} onChange={onChange}>
        {items.map((each) => {
          return (
            <MenuItem value={each} key={each}>
              <InputText>{t[each.toUpperCase()]}</InputText>
            </MenuItem>
          );
        })}
      </SimpleSelectWrapper>
    </>
  );
};

export default OrderByField;
