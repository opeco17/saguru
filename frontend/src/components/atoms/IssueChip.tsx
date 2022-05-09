import Chip from '@mui/material/Chip';
import { useTheme } from '@mui/material/styles';
import { ReactElement } from 'react';

type IssueChipProps = {
  label: string;
  icon?: ReactElement;
};

const IssueChip = ({ label, icon }: IssueChipProps) => {
  const theme = useTheme();
  return (
    <Chip
      label={label}
      icon={icon}
      // @ts-ignore to use custom color
      color='lightGreyChip'
      variant='outlined'
      size='small'
      sx={{
        mr: 1,
        borderRadius: '5px',
        px: 0.5,
        color: theme.palette.lightGreyChip.main,
        fontSize: 13,
      }}
    />
  );
};

export default IssueChip;
