import React, { lazy, Suspense } from 'react';

const LazyPaginator = lazy(() => import('./Paginator'));

const Paginator = (props: JSX.IntrinsicAttributes & { children?: React.ReactNode; }) => (
  <Suspense fallback={null}>
    <LazyPaginator {...props} />
  </Suspense>
);

export default Paginator;
