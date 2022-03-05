import Chip from '@mui/material/Chip';
import { useTheme } from '@mui/material/styles';

type RepositoryChipProps = {
  label: string;
  icon?: any;
};

const RepositoryChip = (props: RepositoryChipProps) => {
  const theme = useTheme();
  const pl = props.icon ? 1 : 0.7;

  return (
    <Chip
      label={props.label}
      icon={props.icon}
      // @ts-ignore to use custom color
      color='info'
      variant='outlined'
      size='small'
      sx={{ mr: 1, borderRadius: '5px', pl: pl, pr: 0.7, color: theme.palette.info.light }}
    />
  );
};

export default RepositoryChip;
