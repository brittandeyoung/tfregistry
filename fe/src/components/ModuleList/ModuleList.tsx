import React, { useEffect,  } from 'react';
import { useParams } from 'react-router-dom';
import { ModuleListWrapper } from './ModuleList.styled';

type TFModule = {
    id: string
    description: string
    downloads: number
    name: string
    namespace: string
    provider: string
    source: string
    verified: boolean
    version: string
}

type ModuleListState = {
    Modules: TFModule[];
}

const ModuleList = () => {
    let { namespace } = useParams();
    const [modules, setModules] = React.useState<ModuleListState>({ "Modules": [] });
    useEffect(() => {
        fetch(process.env.REACT_APP_API_URL_BASE + "/modules/" + namespace)
            .then((response) => response.json())
            .then((data) => {
                setModules(data);
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, [namespace]);

    return (<ModuleListWrapper data-testid="ModuleList">
        <div className="ui left aligned basic padded segment">
            <div className="ui grid">
                <div className="sixteen wide column">
                    <h1>{namespace} / modules</h1>
                    <div className="ui basic segment">
                        <a className="ui primary button"
                            href="/">Create</a>
                    </div>
                    <div className="ui divider"></div>
                    <div className="ui middle aligned list">
                        {modules.Modules?.map((object, i) => (
                            <div className="ui top attached segment" key={i}>
                                <h4 className="header">
                                    {namespace}/{object.name}/{object.provider}
                                    {/* <span className="right floated"></span>
                                {% if  object.moduleversion_set.first.version %}
                                    <a data-tooltip="Latest Version"
                                       className="ui primary button"
                                       href="{% url 'module_version_detail_latest' organization=organization name=module.name provider=module.provider version='latest' %}">
                                        {{ module.moduleversion_set.first.version }}
                                    </a>
                                {% else %}
                                       href="{% url 'module_version_create' organization=organization name=module.name provider=module.provider %}">Create Version</a>
                                <a {% if  module.moduleversion_set.first.version %} href="{% url 'module_version_detail_latest' organization=organization name=module.name provider=module.provider version='latest' %}" {% endif %}> {{ organization }} / {{ module.name }} / {{ module.provider }}</a>
                                <a className="ui right floated primary button"
                                   href="{% url 'module_update' organization=organization name=module.name provider=module.provider %}">Update</a> */}
                                </h4>
                                <div className="meta">
                                    {object.downloads} <i className="download icon"></i>
                                    <span>{object.description}</span>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    </ModuleListWrapper>)
};

export default ModuleList;
