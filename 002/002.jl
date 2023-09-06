dict = Dict()
strs = readline()

for str in strs
    if str ∈ keys(dict)
        dict[str] += 1
    else
        dict[str] = 1
    end
end

dict = sort(collect(dict))

for (key, value) in dict
    println("$(key) => $(value)")
end
