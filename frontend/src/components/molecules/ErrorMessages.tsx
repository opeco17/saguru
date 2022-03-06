import Alert from '@mui/material/Alert';
import Stack from '@mui/material/Stack';

type ErrorMessagesProps = {
  errorMessages: string[];
};

const ErrorMessages = ({ errorMessages }: ErrorMessagesProps) => {
  return (
    <Stack spacing={2} sx={{ mb: 2 }}>
      {errorMessages.map((errorMessage, index) => {
        return (
          <Alert variant='outlined' severity='error' key={index}>
            {errorMessage}
          </Alert>
        );
      })}
    </Stack>
  );
};

export default ErrorMessages;
