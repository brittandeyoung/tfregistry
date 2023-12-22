import React, { lazy, Suspense } from 'react';

const LazyNamespaceList = lazy(() => import('./NamespaceList'));

const NamespaceList = (props: JSX.IntrinsicAttributes & { children?: React.ReactNode; }) => (
  <Suspense fallback={null}>
    <LazyNamespaceList {...props} />
  </Suspense>
);

export default NamespaceList;
