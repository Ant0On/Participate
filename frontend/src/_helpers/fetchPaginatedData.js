import { fetchWrapper } from "@/_helpers/fetch-wrapper";


export default async function* fetchPaginatedData(url, onMap=(data)=>data){
    let page = 1;
    let response = await fetchPage(url, page)

    if(response)
    {
        const totalPages = response?.totalPages;
        yield onMap(response?.data);

        page += 1;

        while(page <= totalPages){
            yield await onMap(fetchPage(url, page)?.data)
            page += 1;
        }
    }
}

async function fetchPage(url, pageNumber){
    return await fetchWrapper.get(`${url}?page=${pageNumber}`)
}