import { Issue } from '../../types/issue';
import IssueChip from '../atoms/IssueChip';
import LabelOutlinedIcon from '@mui/icons-material/LabelOutlined';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { grey } from '@mui/material/colors';
import { useTheme } from '@mui/material/styles';

type IssueCardProps = {
  issue: Issue;
};

const IssueCard = ({ issue }: IssueCardProps) => {
  const theme = useTheme();
  return (
    <Box sx={{ my: 1.5 }}>
      <Box sx={{ display: 'flex', alignItems: 'center', flexWrap: 'wrap', mb: 1.6 }}>
        <Typography
          sx={{
            fontSize: 17,
            mr: 1,
          }}
        >
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
        <Typography
          sx={{
            color: grey['700'],
            fontSize: 14,
          }}
        >
          {`Opend on ${issue.gitHubCreatedAtFormatted} / ${issue.commentCount} comments`}
        </Typography>
      </Box>
      <Box sx={{ display: 'flex', alignItems: 'center', flexWrap: 'wrap', mb: 1.5, rowGap: 1 }}>
        {issue.labels.map((label) => {
          return <IssueChip label={label} key={label} />;
        })}
      </Box>
    </Box>
  );
};

export default IssueCard;
