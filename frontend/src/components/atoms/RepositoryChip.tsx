import Chip from '@mui/material/Chip';
import { useTheme } from '@mui/material/styles';
import { ReactElement } from 'react';

type RepositoryChipProps = {
  label: string;
  icon?: ReactElement;
};

const RepositoryChip = ({ label, icon }: RepositoryChipProps) => {
  const theme = useTheme();
  const pl = icon ? 1 : 0.7;

  return (
    <Chip
      label={label}
      icon={icon}
      // @ts-ignore to use custom color
      color='info'
      variant='outlined'
      size='small'
      sx={{ mr: 1, borderRadius: '5px', pl: pl, pr: 0.7, color: theme.palette.info.light }}
    />
  );
};

export default RepositoryChip;
