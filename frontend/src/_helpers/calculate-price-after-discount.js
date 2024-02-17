function calculatePriceAfterDiscount(price, discount) {
    try{
        if(discount === 0)
            return price
        return price - price * discount / 100
    }
    catch{
        return price
    }

}

export default calculatePriceAfterDiscount