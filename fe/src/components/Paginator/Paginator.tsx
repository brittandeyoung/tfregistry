import React from 'react';
import { PaginatorWrapper } from './Paginator.styled';

type PaginatorProps = {
   next_url?: string;
}

const Paginator = ({ next_url }: PaginatorProps) => {
   if (next_url) {
      return (
         <PaginatorWrapper data-testid="Paginator">
            <div className="ui footer segment">
               <div className="ui container">
                  <div className="ui pagination menu">
                     <a className="item" href={next_url}><i className="angle right icon"></i></a>
                  </div>
               </div>
            </div>

         </PaginatorWrapper>
      );
   } else {
      return (<div></div>);
   }
}

export default Paginator;
