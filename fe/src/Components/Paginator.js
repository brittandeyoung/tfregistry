import React, { useState } from "react";

export default function Paginator ({next_url}) {
    if (next_url) {
  return (
    <div class="ui footer segment">
    <div class="ui container">
        <div class="ui pagination menu">
            <a class="item" href={next_url}><i class="angle right icon"></i></a>
        </div>
    </div>
</div>
  );
    } else {
        return (<div></div>)
    }
};
