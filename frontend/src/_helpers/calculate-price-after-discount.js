function calculatePriceAfterDiscount(price, discount) {
    try{
        console.log(price, discount)
        if(discount === 0)
            return price
        return price - price * discount / 100
    }
    catch{
        return price
    }

}

export default calculatePriceAfterDiscount