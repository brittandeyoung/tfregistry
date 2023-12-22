import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import NamespaceList from './NamespaceList';

describe('<NamespaceList />', () => {
  test('it should mount', () => {
    render(<NamespaceList />);
    
    const namespaceList = screen.getByTestId('NamespaceList');

    expect(namespaceList).toBeInTheDocument();
  });
});