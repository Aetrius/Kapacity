import * as React from 'react';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Divider,
  IconButton,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography
} from '@mui/material';
// import { ArrowRight as ArrowRightIcon } from '@mui/icons-material';
// import { ContactlessPayment as ContactlessPaymentIcon } from '@mui/icons-material';
import { DotsThree as DotsThreeIcon } from '@phosphor-icons/react/dist/ssr/DotsThree';
import { Application } from './application';

export interface ApplicationProps {
  applications: Application[];
}

const headers = ['Name', 'Cluster', 'Memory Value', 'Memory Request', 'Memory Limit', 'CPU Value', 'CPU Request', 'CPU Limit'];

export function ApplicationTable({ applications }: ApplicationProps): React.JSX.Element {
  return (
    <Card>
      <CardHeader
        title="Applications"
      />
      <CardContent sx={{ pb: '8px' }}>
        <Box sx={{ minWidth: 800 }}>
          <Table>
            <TableHead>
              <TableRow>
                {headers.map((header) => (
                  <TableCell key={header}>
                    <Typography variant="subtitle2" fontWeight="bold">
                      {header}
                    </Typography>
                  </TableCell>
                ))}
                <TableCell />
              </TableRow>
            </TableHead>
            <TableBody>
              {applications.map((application) => (
                <TableRow hover key={application.id}>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography noWrap variant="subtitle2">
                        {application.title}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.cluster}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.memoryValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.memoryRequest}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.memoryLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.cpuValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.cpuRequest}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{application.cpuLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <IconButton>
                      <DotsThreeIcon weight="bold" />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Box>
      </CardContent>
      <Divider />
      <CardActions>
        <Button color="secondary" endIcon={<DotsThreeIcon />} size="small">
          See all applications
        </Button>
      </CardActions>
    </Card>
  );
}
