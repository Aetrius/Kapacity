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
import { Node } from './node';

export interface NodeProps {
  nodes: Node[];
}

const headers = ['Name', 'Cluster', 'Memory Value', 'Memory Limit', 'CPU Value', 'CPU Limit', 'Disk Value', 'Disk Size'];

export function NodeTable({ nodes }: NodeProps): React.JSX.Element {
  return (
    <Card>
      <CardHeader
        title="Nodes"
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
              {nodes.map((node) => (
                <TableRow hover key={node.id}>
                  <TableCell>
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography noWrap variant="subtitle2">
                        {node.name}
                      </Typography>
                    </Box>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.cluster}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.memoryValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.memoryLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.cpuValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.cpuLimit}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.diskValue}</Typography>
                  </TableCell>
                  <TableCell>
                    <Typography variant="body2">{node.diskSize}</Typography>
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
          See all nodes
        </Button>
      </CardActions>
    </Card>
  );
}
