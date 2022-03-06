import { Issue } from '../../types/issue';
import IssueChip from '../atoms/IssueChip';
import LabelOutlinedIcon from '@mui/icons-material/LabelOutlined';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';

type IssueCardProps = {
  issue: Issue;
};

const IssueCard = ({ issue }: IssueCardProps) => {
  const theme = useTheme();
  return (
    <Box sx={{ display: 'flex', alignItems: 'center', flexWrap: 'wrap', my: 1.5 }}>
      {/* @ts-ignore to use custom color */}
      <LabelOutlinedIcon size='small' sx={{ color: 'text.secondary', mr: 1.5 }} />
      <Typography sx={{ mr: 1.5 }}>
        <Box
          component='a'
          href={issue.url}
          target='_blank'
          rel='noreferrer'
          sx={{ '&:hover': { color: theme.palette.secondary.main } }}
        >
          {issue.title}
        </Box>
      </Typography>
      {issue.labels.map((label) => {
        return <IssueChip label={label} key={label} />;
      })}
    </Box>
  );
};

export default IssueCard;
