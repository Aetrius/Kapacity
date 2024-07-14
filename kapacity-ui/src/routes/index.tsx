import * as React from 'react';
import type { RouteObject } from 'react-router-dom';
import { Outlet } from 'react-router-dom';

import { Page as HomePage } from '@/pages/marketing/home';
import { Page as Overview } from '@/pages/dashboard/overview';
import { Clusters } from '@/pages/dashboard/clusters';
import { Applications } from '@/pages/dashboard/applications';
import { Storage } from '@/pages/dashboard/storage';
import { Nodes } from '@/pages/dashboard/nodes';
import { Page as NotFoundPage } from '@/pages/not-found';
import { Layout as DashboardLayout } from '@/components/dashboard/layout/layout';

import { route as authRoute } from './auth';
import { route as dashboardRoute } from './dashboard';


export const routes: RouteObject[] = [
  {
    element: (
      <DashboardLayout>
        <Outlet />
      </DashboardLayout>
    ),
    children: [
      { index: true, element: <Overview /> },
      { path: 'overview', element: <Overview /> },
      { path: 'clusters', element: <Clusters /> },
      { path: 'applications', element: <Applications /> },
      { path: 'storage', element: <Storage /> },
      { path: 'nodes', element: <Nodes /> },
    ],
  },
  authRoute,
  dashboardRoute,
  { path: '*', element: <NotFoundPage /> },
];
