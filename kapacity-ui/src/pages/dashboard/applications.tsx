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
import { Application } from '@/components/dashboard/overview/application';
import { Dropdown } from '@/components/core/dropdown/dropdown';
import { DropdownTrigger } from '@/components/core/dropdown/dropdown-trigger';
import { DropdownPopover } from '@/components/core/dropdown/dropdown-popover';
import MenuList from '@mui/material/MenuList';
import MenuItem from '@mui/material/MenuItem';


const metadata = { title: `Overview | Applications | ${config.site.name}` } satisfies Metadata;

const apps: Application[] = [
  {
    id: '1',
    title: 'App One',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryRequest: '14Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuRequest: '3000m',
    cpuLimit: '5000m',
  },
  {
    id: '12',
    title: 'App Two',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryRequest: '14Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuRequest: '3000m',
    cpuLimit: '5000m',
  },
  {
    id: '12',
    title: 'App Three',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryRequest: '14Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuRequest: '3000m',
    cpuLimit: '5000m',
  },
  {
    id: '12',
    title: 'App Four',
    icon: 'icon1.png',
    cluster: 'cluster1',
    memoryValue: '12Gib',
    memoryRequest: '14Gib',
    memoryLimit: '20Gib',
    cpuValue: '2000m',
    cpuRequest: '3000m',
    cpuLimit: '5000m',
  },
];

export function Applications(): React.JSX.Element {
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
              <Typography variant="h4">Applications</Typography>
            </Box>
            <Box>
              <Dropdown>
                <DropdownTrigger>
                  <Button variant="contained">Owner: All</Button>
                </DropdownTrigger>
                <DropdownPopover>
                  <MenuList>
                    <MenuItem>Owner A</MenuItem>
                    <MenuItem>Owner B</MenuItem>
                    <MenuItem>Owner C</MenuItem>
                  </MenuList>
                </DropdownPopover>
              </Dropdown>
            </Box>
            <Box>
              <Dropdown>
                <DropdownTrigger>
                  <Button variant="contained">Application: All</Button>
                </DropdownTrigger>
                <DropdownPopover>
                  <MenuList>
                    <MenuItem>Application A</MenuItem>
                    <MenuItem>Application B</MenuItem>
                    <MenuItem>Application C</MenuItem>
                  </MenuList>
                </DropdownPopover>
              </Dropdown>
            </Box>
          </Stack>
          <Grid container spacing={4}>
            <Grid md={12} xs={12}>
              <ApplicationTable  applications={apps}/>
            </Grid>
          </Grid>
        </Stack>
      </Box>
    </React.Fragment>
  );
}
