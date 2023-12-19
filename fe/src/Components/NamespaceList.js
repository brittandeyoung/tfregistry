import React, { useEffect, useState } from "react";
import Paginator from "./Paginator"; 

const NamespaceList = () => {

    const [namespaces, setNamespaces] = useState([]);
   useEffect(() => {
      fetch(process.env.REACT_APP_API_URL_BASE + "/namespaces")
         .then((response) => response.json())
         .then((data) => {
            console.log(data);
            setNamespaces(data);
         })
         .catch((err) => {
            console.log(err.message);
         });
   }, []);

    return (
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
                    {namespaces.Namespaces?.map((object, i) => (
                    <div className="card" key={i}>
                        <div className="content">
                            <div className="header">
                                <a href={object.name}>{object.name}</a>
                            </div>
                            <div className="description">{object.description}</div>
                        </div>
                    </div>))}
            </div>
        </div>
        </div>
        </div>
        <Paginator next_url={namespaces.Meta?.next_url} />
        </div>
        );
};

export default NamespaceList;