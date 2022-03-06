import Chip from '@mui/material/Chip';
import { useTheme } from '@mui/material/styles';

type IssueChipProps = {
  label: string;
};

const IssueChip = ({ label }: IssueChipProps) => {
  const theme = useTheme();
  return (
    <Chip
      label={label}
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
