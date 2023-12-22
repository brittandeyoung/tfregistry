import React, { lazy, Suspense } from 'react';

const LazyModuleList = lazy(() => import('./ModuleList'));

const ModuleList = (props: JSX.IntrinsicAttributes & { children?: React.ReactNode; }) => (
  <Suspense fallback={null}>
    <LazyModuleList {...props} />
  </Suspense>
);

export default ModuleList;
