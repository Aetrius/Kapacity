import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';
import { ArrowRight as ArrowRightIcon } from '@phosphor-icons/react/dist/ssr/ArrowRight';
import { Briefcase as BriefcaseIcon } from '@phosphor-icons/react/dist/ssr/Briefcase';
import { FileCode as FileCodeIcon } from '@phosphor-icons/react/dist/ssr/FileCode';
import { Info as InfoIcon } from '@phosphor-icons/react/dist/ssr/Info';
import { ListChecks as ListChecksIcon } from '@phosphor-icons/react/dist/ssr/ListChecks';
import { Plus as PlusIcon } from '@phosphor-icons/react/dist/ssr/Plus';
import { Users as UsersIcon } from '@phosphor-icons/react/dist/ssr/Users';
import { Warning as WarningIcon } from '@phosphor-icons/react/dist/ssr/Warning';
import { Helmet } from 'react-helmet-async';

import type { Metadata } from '@/types/metadata';
import { config } from '@/config';
import { dayjs } from '@/lib/dayjs';
import { AppChat } from '@/components/dashboard/overview/app-chat';
import { AppLimits } from '@/components/dashboard/overview/app-limits';
import { AppUsage } from '@/components/dashboard/overview/app-usage';
import { Events } from '@/components/dashboard/overview/events';
import { HelperWidget } from '@/components/dashboard/overview/helper-widget';
import { Subscriptions } from '@/components/dashboard/overview/subscriptions';
import { Summary } from '@/components/dashboard/overview/summary';
import { SummaryPercent } from '@/components/dashboard/overview/summaryPerecent';
import { ApplicationTable } from '@/components/dashboard/overview/appTable';
import { Node } from '@/components/dashboard/overview/node';
import { NodeTable } from '@/components/dashboard/overview/nodeTable';
import { Dropdown } from '@/components/core/dropdown/dropdown';
import { DropdownTrigger } from '@/components/core/dropdown/dropdown-trigger';
import { DropdownPopover } from '@/components/core/dropdown/dropdown-popover';
import MenuList from '@mui/material/MenuList';
import MenuItem from '@mui/material/MenuItem';

const metadata = { title: `Overview | Nodes | ${config.site.name}` } satisfies Metadata;

const nodes: Node[] = [
  {
    id: '1',
    name: 'Node One',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    diskValue: '50Gib',
    diskSize: '100Gib',
  },
  {
    id: '11',
    name: 'Node Two',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    diskValue: '50Gib',
    diskSize: '100Gib',
  },
  {
    id: '12',
    name: 'Node Three',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    diskValue: '50Gib',
    diskSize: '100Gib',
  },
  {
    id: '13',
    name: 'Node Four',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuLimit: '5000m',
    diskValue: '50Gib',
    diskSize: '100Gib',
  },
];

export function Nodes(): React.JSX.Element {
  return (
    <React.Fragment>
      <Helmet>
        <title>{metadata.title}</title>
      </Helmet>
      <Box
        sx={{
          maxWidth: 'var(--Content-maxWidth)',
          m: 'var(--Content-margin)',
          p: 'var(--Content-padding)',
          width: 'var(--Content-width)',
        }}
      >
        <Stack spacing={4}>
          <Stack direction={{ xs: 'column', sm: 'row' }} spacing={3} sx={{ alignItems: 'flex-start' }}>
            <Box sx={{ flex: '1 1 auto' }}>
              <Typography variant="h4">Nodes</Typography>
            </Box>
            <Box>
            <Dropdown>
              <DropdownTrigger>
              <Button variant="contained">Environment: All</Button>
              </DropdownTrigger>
              <DropdownPopover>
                <MenuList>
                  <MenuItem>PROD</MenuItem>
                  <MenuItem>STAGE</MenuItem>
                  <MenuItem>QA</MenuItem>
                </MenuList>
              </DropdownPopover>
            </Dropdown>
            </Box>
            <Box>
            <Dropdown>
              <DropdownTrigger>
              <Button variant="contained">Cluster: All</Button>
              </DropdownTrigger>
              <DropdownPopover>
                <MenuList>
                  <MenuItem>Cluster 1</MenuItem>
                  <MenuItem>Cluster 2</MenuItem>
                  <MenuItem>Cluster 3</MenuItem>
                </MenuList>
              </DropdownPopover>
            </Dropdown>
            </Box>
          </Stack>
          <Grid container spacing={4}>
            <Grid md={12} xs={12}>
              <NodeTable nodes={nodes}/>
            </Grid>
          </Grid>
          
        </Stack>
      </Box>
    </React.Fragment>
  );
}
