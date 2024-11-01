import { useState, useEffect } from 'react';
import { NamespaceListWrapper } from './NamespaceList.styled';
import { useSearchParams } from "react-router-dom";

type Namespace = {
   name: string;
   description: string;
}

const NamespaceList = () => {
   const [searchParams] = useSearchParams();

   const limit = searchParams.get('limit') || '10'; 
   const [items, setItems] = useState<Namespace[]>([]);
   const [page, setPage] = useState(1);
   const [startKey, setStartKey] = useState("")

   async function fetchMoreData(limit: string, startKey: string) {
      const response = await fetch(process.env.REACT_APP_API_URL_BASE + "/api/namespaces?limit=" + limit + (startKey ? "&start_key=" + startKey : ""));
      const newData = await response.json();
      return newData;
    }

   const fetchData = async () => {
      const newData = await fetchMoreData(limit, startKey);
      const lek = newData.Meta.last_evaluated_key
      setPage(page + 1);
      setStartKey(lek)
      setItems([...items, ...newData.Namespaces]);
   };

   useEffect(() => {
      fetchData(); // Fetch initial data on component mount
    }, []); // eslint-disable-line react-hooks/exhaustive-deps

   return (<NamespaceListWrapper data-testid="NamespaceList">
      <div className="ui left aligned basic padded segment">
         <div className="ui grid">
            <div className="sixteen wide column">
               <h1>Namespaces</h1>
               <div className="ui basic segment">
                  <a className="ui primary button" href="/create">Create</a>
                  <div className="ui icon input">
                     {/* <form id="desktop_search"
                      method="GET"
                      action="{% url 'organization_list' %}"
                      accept-charset="utf-8">
                </form> */}
                     {/* <input name="query"
                       type="text"
                       placeholder="Search by Name"
                       value="{{ request.GET.query }}"
                       form="desktop_search"> </input> */}
                  </div>
                  <div className="ui divider"></div>
                  <div className="ui cards">
                  {items.map(item => ( (
                        <div className="card" key={item.name}>
                           <div className="content">
                              <div className="header">
                                 <a href={item.name}>{item.name}</a>
                              </div>
                              <div className="description">{item.description}</div>
                           </div>
                        </div>)))}
                  </div>
               </div>
            </div>
         </div>
         <div className="ui footer segment">
            {startKey === null && "All caught up!" }
            {startKey !== null && <button className="fluid ui button" onClick={fetchData}>
               Load More..
            </button>}
         </div>
      </div>
   </NamespaceListWrapper>);
};

export default NamespaceList;
