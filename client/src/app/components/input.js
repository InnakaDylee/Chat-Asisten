export function Input({input, onChange}) {
  return (
    <input spellCheck={"false"} className="text-white w-[85%] p-3 border-solid bg-gray-600 bg-opacity-40 rounded-[30px]" type="text" value={input} placeholder="Ketik Pesan..." onChange={onChange}/>
  )
}